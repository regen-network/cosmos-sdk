package simapp

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dist036 "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v0_36"
	gov036 "github.com/cosmos/cosmos-sdk/x/gov/legacy/v036"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (app *SimApp) inplaceMigration(chainID string) error {
	now := time.Now().UTC()
	header := abci.Header{ChainID: chainID, Time: now}
	ctx := app.GetDeliverState(header)

	return InplaceMigration034to036(ctx, app.keyGov, app.keyDistr)
}

// InplaceMigration034to036 migrates a 0.34 datastore to 0.36 format
// This can be registered as a handler for the x/upgrade module
// (once that is merged in a separate PR)
// Should be done in the app setup, eg.
//
//   app.upgradeKeeper.SetUpgradeHandler("v0_36", func(ctx sdk.Context, plan upgrade.Plan) {
//     InplaceMigrationFromv034(ctx, app.govKey, app.distrKey)
//   })
func InplaceMigration034to036(ctx sdk.Context, govKey sdk.StoreKey, distKey sdk.StoreKey) error {
	err := dist036.InplaceMigrationFromv034(ctx, distKey)
	if err != nil {
		return err
	}
	err = gov036.InplaceMigrationFromv034(ctx, govKey)
	return err
}
