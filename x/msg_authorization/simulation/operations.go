package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/keeper"
	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Simulation operation weights constants
const (
	OpWeightMsgGrantAuthorization  = "op_weight_msg_grant_authorization"
	OpWeightMsgRevokeAuthorization = "op_weight_msg_revoke_authorization"
	OpWeightMsgExecAuthorization   = "op_weight_msg_exec_authorization"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc *codec.Codec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgGrantAuthorization  int
		weightMsgRevokeAuthorization int
		weightMsgExecAuthorization   int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgGrantAuthorization, &weightMsgGrantAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgGrantAuthorization = simappparams.DefaultWeightMsgGrantAuthorization
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRevokeAuthorization, &weightMsgRevokeAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeAuthorization = simappparams.DefaultWeightMsgRevokeAuthorization
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgExecAuthorization, &weightMsgExecAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgExecAuthorization = simappparams.DefaultWeightMsgExecAuthorization
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgGrantAuthorization,
			SimulateMsgGrantAuthorization(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRevokeAuthorization,
			SimulateMsgRevokeAuthorization(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgExecAuthorization,
			SimulateMsgExecAuthorization(ak, bk, k),
		),
	}
}

// SimulateMsgGrantAuthorization generates a MsgGrantAuthorization with random values
// nolint: interfacer
func SimulateMsgGrantAuthorization(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account,
		chainID string) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		granterAcc, _ := simtypes.RandomAcc(r, accounts)
		granteeAcc, _ := simtypes.RandomAcc(r, accounts)
		authorization, expiration := k.GetAuthorization(ctx, granteeAcc.Address, granterAcc.Address, bank.MsgSend{}.Type())

		if authorization != nil && expiration != 0 {
			return simtypes.NoOpMsg(types.ModuleName, "", ""), nil, nil
		}

		now := ctx.BlockHeader().Time
		newCoins := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 100))
		x := &types.SendAuthorization{SpendLimit: newCoins}
		msgGrantAuthorization, err := types.NewMsgGrantAuthorization(granteeAcc.Address, granterAcc.Address, x, now.Add(3600))

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgGrantAuthorization.Type(), ""), nil, nil
		}

		var fees sdk.Coins
		fees, err = simtypes.RandomFees(r, ctx, newCoins)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgGrantAuthorization.Type(), ""), nil, err
		}

		granterAccount := ak.GetAccount(ctx, granterAcc.Address)

		tx := helpers.GenTx(
			[]sdk.Msg{msgGrantAuthorization},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{granterAccount.GetAccountNumber()},
			[]uint64{granterAccount.GetSequence()},
			granterAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgGrantAuthorization.Type(), ""), nil, err
		}

		return simtypes.NewOperationMsg(msgGrantAuthorization, true, ""), nil, nil
	}
}

// SimulateMsgRevokeAuthorization generates a MsgRevokeAuthorization with random values
// nolint: interfacer
func SimulateMsgRevokeAuthorization(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account,
		chainID string) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		granterAcc, _ := simtypes.RandomAcc(r, accounts)
		granteeAcc, _ := simtypes.RandomAcc(r, accounts)
		authorization, expiration := k.GetAuthorization(ctx, granteeAcc.Address, granterAcc.Address, bank.MsgSend{}.Type())

		if authorization != nil || expiration != 0 {
			return simtypes.NoOpMsg(types.ModuleName, "", ""), nil, nil
		}

		now := ctx.BlockHeader().Time
		newCoins := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 100))
		x := &types.SendAuthorization{SpendLimit: newCoins}
		xAny, err := types.ConvertToAny(x)

		k.Grant(ctx, granteeAcc.Address, granterAcc.Address, xAny, now.Unix()+3600)
		msgRevokeAuthorization := types.NewMsgRevokeAuthorization(granterAcc.Address, granteeAcc.Address, bank.MsgSend{}.Type())

		var fees sdk.Coins
		granterAccount := ak.GetAccount(ctx, granterAcc.Address)
		spendable := bk.SpendableCoins(ctx, granterAccount.GetAddress())
		fees, err = simtypes.RandomFees(r, ctx, spendable)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgRevokeAuthorization.Type(), ""), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msgRevokeAuthorization},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{granterAccount.GetAccountNumber()},
			[]uint64{granterAccount.GetSequence()},
			granterAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgRevokeAuthorization.Type(), ""), nil, err
		}

		return simtypes.NewOperationMsg(msgRevokeAuthorization, true, ""), nil, nil
	}
}

// SimulateMsgExecAuthorization generates a MsgExecAuthorization with random values
// nolint: interfacer
func SimulateMsgExecAuthorization(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account,
		chainID string) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		granterAcc, _ := simtypes.RandomAcc(r, accounts)
		granteeAcc, _ := simtypes.RandomAcc(r, accounts)
		recipientAcc, _ := simtypes.RandomAcc(r, accounts)
		denom := sdk.DefaultBondDenom
		authorization, expiration := k.GetAuthorization(ctx, granteeAcc.Address, granterAcc.Address, bank.MsgSend{}.Type())

		if authorization != nil || expiration != 0 {
			return simtypes.NoOpMsg(types.ModuleName, "", ""), nil, nil
		}

		now := ctx.BlockHeader().Time
		newCoins := sdk.NewCoins(sdk.NewInt64Coin(denom, 100))
		x := &types.SendAuthorization{SpendLimit: newCoins}
		xAny, err := types.ConvertToAny(x)

		k.Grant(ctx, granteeAcc.Address, granterAcc.Address, xAny, now.Unix()+3600)
		authorization, expiration = k.GetAuthorization(ctx, granteeAcc.Address, granterAcc.Address, bank.MsgSend{}.Type())

		if authorization == nil || expiration == 0 {
			return simtypes.NoOpMsg(types.ModuleName, "", ""), nil, nil
		}

		msgs := []sdk.Msg{
			&bank.MsgSend{
				Amount:      sdk.NewCoins(sdk.NewInt64Coin(denom, 2)),
				FromAddress: granterAcc.Address,
				ToAddress:   recipientAcc.Address,
			},
		}

		msgExecAuthorization, err := types.NewMsgExecAuthorized(granteeAcc.Address, msgs)

		var fees sdk.Coins
		granteeAccount := ak.GetAccount(ctx, granteeAcc.Address)
		spendable := bk.SpendableCoins(ctx, granteeAccount.GetAddress())
		fees, err = simtypes.RandomFees(r, ctx, spendable)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgExecAuthorization.Type(), ""), nil, err
		}

		tx := helpers.GenTx(
			[]sdk.Msg{&msgExecAuthorization},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{granteeAccount.GetAccountNumber()},
			[]uint64{granteeAccount.GetSequence()},
			granteeAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msgExecAuthorization.Type(), ""), nil, err
		}

		return simtypes.NewOperationMsg(&msgExecAuthorization, true, ""), nil, nil
	}
}
