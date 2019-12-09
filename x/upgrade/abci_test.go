package upgrade_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
)

type TestSuite struct {
	suite.Suite

	module                module.AppModule
	keeper                upgrade.Keeper
	querier               sdk.Querier
	handler               gov.Handler
	ctx                   sdk.Context
	FlagUnsafeSkipUpgrade string
}

func (s *TestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)

	s.keeper = app.UpgradeKeeper
	s.handler = upgrade.NewSoftwareUpgradeProposalHandler(s.keeper)
	s.querier = upgrade.NewQuerier(s.keeper)
	s.ctx = app.BaseApp.NewContext(checkTx, abci.Header{Height: 10, Time: time.Now()})
	s.module = upgrade.NewAppModule(s.keeper)
	s.FlagUnsafeSkipUpgrade = upgrade.FlagUnsafeSkipUpgrade
	viper.Set(upgrade.FlagUnsafeSkipUpgrade, []int{})
	s.VerifySet()
}

func (s *TestSuite) TestRequireName() {
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{}})
	s.Require().NotNil(err)
	s.Require().Equal(sdk.CodeUnknownRequest, err.Code())
}

func (s *TestSuite) TestRequireFutureTime() {
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: s.ctx.BlockHeader().Time}})
	s.Require().NotNil(err)
	s.Require().Equal(sdk.CodeUnknownRequest, err.Code())
}

func (s *TestSuite) TestRequireFutureBlock() {
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight()}})
	s.Require().NotNil(err)
	s.Require().Equal(sdk.CodeUnknownRequest, err.Code())
}

func (s *TestSuite) TestCantSetBothTimeAndHeight() {
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now(), Height: s.ctx.BlockHeight() + 1}})
	s.Require().NotNil(err)
	s.Require().Equal(sdk.CodeUnknownRequest, err.Code())
}

func (s *TestSuite) TestDoTimeUpgrade() {
	s.T().Log("Verify can schedule an upgrade")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	s.Require().Nil(err)

	s.VerifyDoUpgrade()
}

func (s *TestSuite) TestDoHeightUpgrade() {
	s.T().Log("Verify can schedule an upgrade")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)

	s.VerifyDoUpgrade()
}

func (s *TestSuite) TestCanOverwriteScheduleUpgrade() {
	s.T().Log("Can overwrite plan")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "bad_test", Height: s.ctx.BlockHeight() + 10}})
	s.Require().Nil(err)
	err = s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)

	s.VerifyDoUpgrade()
}

func (s *TestSuite) VerifyDoUpgrade() {
	s.T().Log("Verify that a panic happens at the upgrade time/height")
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	s.Require().Panics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.T().Log("Verify that the upgrade can be successfully applied with a handler")
	s.keeper.SetUpgradeHandler("test", func(ctx sdk.Context, plan upgrade.Plan) {})
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.VerifyCleared(newCtx)
}

func (s *TestSuite) VerifyDoUpgradeWithCtx(newCtx sdk.Context, proposalName string) {
	s.T().Log("Verify that a panic happens at the upgrade time/height")
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	s.Require().Panics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.T().Log("Verify that the upgrade can be successfully applied with a handler")
	s.keeper.SetUpgradeHandler(proposalName, func(ctx sdk.Context, plan upgrade.Plan) {})
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.VerifyCleared(newCtx)
}

func (s *TestSuite) TestHaltIfTooNew() {
	s.T().Log("Verify that we don't panic with registered plan not in database at all")
	var called int
	s.keeper.SetUpgradeHandler("future", func(ctx sdk.Context, plan upgrade.Plan) { called++ })

	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})
	s.Require().Equal(0, called)

	s.T().Log("Verify we panic if we have a registered handler ahead of time")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "future", Height: s.ctx.BlockHeight() + 3}})
	s.Require().NoError(err)
	s.Require().Panics(func() {
		s.module.BeginBlock(newCtx, req)
	})
	s.Require().Equal(0, called)

	s.T().Log("Verify we no longer panic if the plan is on time")

	futCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 3).WithBlockTime(time.Now())
	req = abci.RequestBeginBlock{Header: futCtx.BlockHeader()}
	s.Require().NotPanics(func() {
		s.module.BeginBlock(futCtx, req)
	})
	s.Require().Equal(1, called)

	s.VerifyCleared(futCtx)
}

func (s *TestSuite) VerifyCleared(newCtx sdk.Context) {
	s.T().Log("Verify that the upgrade plan has been cleared")
	bz, err := s.querier(newCtx, []string{upgrade.QueryCurrent}, abci.RequestQuery{})
	s.Require().NoError(err)
	s.Require().Nil(bz)
}

func (s *TestSuite) TestCanClear() {
	s.T().Log("Verify upgrade is scheduled")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	s.Require().Nil(err)

	err = s.handler(s.ctx, upgrade.CancelSoftwareUpgradeProposal{Title: "cancel"})
	s.Require().Nil(err)

	s.VerifyCleared(s.ctx)
}

func (s *TestSuite) TestCantApplySameUpgradeTwice() {
	s.TestDoTimeUpgrade()
	s.T().Log("Verify an upgrade named \"test\" can't be scheduled twice")
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	s.Require().NotNil(err)
	s.Require().Equal(sdk.CodeUnknownRequest, err.Code())
}

func (s *TestSuite) TestNoSpuriousUpgrades() {
	s.T().Log("Verify that no upgrade panic is triggered in the BeginBlocker when we haven't scheduled an upgrade")
	req := abci.RequestBeginBlock{Header: s.ctx.BlockHeader()}
	s.Require().NotPanics(func() {
		s.module.BeginBlock(s.ctx, req)
	})
}

func (s *TestSuite) TestPlanStringer() {
	t, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	s.Require().Nil(err)
	s.Require().Equal(`Upgrade Plan
  Name: test
  Time: 2020-01-01T00:00:00Z
  Info: `, upgrade.Plan{Name: "test", Time: t}.String())
	s.Require().Equal(`Upgrade Plan
  Name: test
  Height: 100
  Info: `, upgrade.Plan{Name: "test", Height: 100}.String())
}

func (s *TestSuite) VerifyNotDone(newCtx sdk.Context, name string) {
	s.T().Log("Verify that upgrade was not done")
	height := s.keeper.GetDoneHeight(newCtx, name)
	s.Require().Zero(height)
}

func (s *TestSuite) VerifyDone(newCtx sdk.Context, name string) {
	s.T().Log("Verify that the upgrade plan has been executed")
	height := s.keeper.GetDoneHeight(newCtx, name)
	s.Require().NotZero(height)
}

func (s *TestSuite) VerifySet() {
	s.T().Log("Verify if the skip upgrade has been set")
	s.Require().NotNil(viper.GetIntSlice(s.FlagUnsafeSkipUpgrade))
}

func (s *TestSuite) VerifyConversion(skipUpgrade []int) {
	skipUpgradeHeights := upgrade.ConvertIntArrayToInt64(skipUpgrade)
	s.Require().Equal(reflect.TypeOf(skipUpgradeHeights).Elem().Kind(), reflect.Int64)
}

func (s *TestSuite) TestContains() {
	viper.Set(s.FlagUnsafeSkipUpgrade, []int{1, 2})
	s.VerifySet()
	skipUpgradeHeights := upgrade.ConvertIntArrayToInt64(viper.GetIntSlice(s.FlagUnsafeSkipUpgrade))
	s.T().Log("case where array contains the element")
	present := upgrade.Contains(skipUpgradeHeights, 1)
	s.Require().True(present)

	s.T().Log("case where array doesn't contain the element")
	present = upgrade.Contains(skipUpgradeHeights, 4)
	s.Require().False(present)
}

func (s *TestSuite) TestSkipUpgradeSkippingAll() {
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)

	s.T().Log("Verify if skip upgrade flag clears upgrade plan in both cases")
	viper.Set(s.FlagUnsafeSkipUpgrade, []int{int(s.ctx.BlockHeight() + 1), int(s.ctx.BlockHeight() + 10)})
	s.VerifySet()

	s.VerifyConversion(viper.GetIntSlice(s.FlagUnsafeSkipUpgrade))

	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.T().Log("Verify a second proposal also is being cleared")
	err = s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: s.ctx.BlockHeight() + 10}})
	s.Require().Nil(err)

	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 10)
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	//To ensure verification is being done only after both upgrades are cleared
	s.T().Log("Verify if both proposals are cleared")
	s.VerifyCleared(s.ctx)
	s.VerifyNotDone(s.ctx, "test")
	s.VerifyNotDone(s.ctx, "test2")
}

func (s *TestSuite) TestUpgradeSkippingOne() {
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)

	s.T().Log("Verify if skip upgrade flag clears upgrade plan in one case and does upgrade on another")
	viper.Set(s.FlagUnsafeSkipUpgrade, []int{int(s.ctx.BlockHeight() + 1)})
	s.VerifySet()

	s.VerifyConversion(viper.GetIntSlice(s.FlagUnsafeSkipUpgrade))
	//Setting block height of proposal test
	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.T().Log("Verify the second proposal is not skipped")
	err = s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: s.ctx.BlockHeight() + 10}})
	s.Require().Nil(err)
	//Setting block height of proposal test2
	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 10)
	s.VerifyDoUpgradeWithCtx(newCtx, "test2")

	s.T().Log("Verify first proposal is cleared and second is done")
	s.VerifyNotDone(s.ctx, "test")
	s.VerifyDone(s.ctx, "test2")
}

func (s *TestSuite) TestUpgradeSkippingOnlyTwo() {
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)

	s.T().Log("Verify if skip upgrade flag clears upgrade plan in both cases and does third upgrade")
	viper.Set(s.FlagUnsafeSkipUpgrade, []int{int(s.ctx.BlockHeight() + 1), int(s.ctx.BlockHeight() + 10)})
	s.VerifySet()

	s.VerifyConversion(viper.GetIntSlice(s.FlagUnsafeSkipUpgrade))

	//Setting block height of proposal test
	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	//A new proposal with height in skipUpgradeHeights
	err = s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: s.ctx.BlockHeight() + 10}})
	s.Require().Nil(err)
	//Setting block height of proposal test2
	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 10)
	s.Require().NotPanics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.T().Log("Verify a new proposal is not skipped")
	err = s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop3", Plan: upgrade.Plan{Name: "test3", Height: s.ctx.BlockHeight() + 15}})
	s.Require().Nil(err)
	newCtx = newCtx.WithBlockHeight(s.ctx.BlockHeight() + 15)
	s.VerifyDoUpgradeWithCtx(newCtx, "test3")

	s.T().Log("Verify two proposals are cleared and third is done")
	s.VerifyNotDone(s.ctx, "test")
	s.VerifyNotDone(s.ctx, "test2")
	s.VerifyDone(s.ctx, "test3")
}

func (s *TestSuite) TestUpgradeWithoutSkip() {
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	s.Require().Nil(err)
	s.T().Log("Verify if upgrade happens without skip upgrade")
	s.Require().Panics(func() {
		s.module.BeginBlock(newCtx, req)
	})

	s.VerifyDoUpgrade()
	s.VerifyDone(s.ctx, "test")
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
