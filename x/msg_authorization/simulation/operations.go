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
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc *codec.Codec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgGrantAuthorization  int
		weightMsgRevokeAuthorization int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgGrantAuthorization, &weightMsgGrantAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgGrantAuthorization = simappparams.DefaultWeightMsgCreateValidator
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRevokeAuthorization, &weightMsgRevokeAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeAuthorization = simappparams.DefaultWeightMsgEditValidator
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
	}
}

// SimulateMsgGrantAuthorization generates a MsgGrantAuthorization with random values
// nolint: interfacer
func SimulateMsgGrantAuthorization(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		granterAcc, _ := simtypes.RandomAcc(r, accounts)
		granteeAcc, _ := simtypes.RandomAcc(r, accounts)

		authorization, expiration := k.GetAuthorization(ctx, granteeAcc.Address, granterAcc.Address, bank.MsgSend{}.Type())

		if authorization == nil && expiration == 0 {
			return simtypes.NoOpMsg(types.ModuleName), nil, nil
		}
		now := ctx.BlockHeader().Time

		newCoins := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 100))
		x := &types.SendAuthorization{SpendLimit: newCoins}

		msgGrantAuthorization, err := types.NewMsgGrantAuthorization(granteeAcc.Address, granterAcc.Address, x, now.Add(3600))
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName), nil, nil
		}
		var fees sdk.Coins
		fees, err = simtypes.RandomFees(r, ctx, newCoins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName), nil, err
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
			return simtypes.NoOpMsg(types.ModuleName), nil, err
		}

		return simtypes.NewOperationMsg(msgGrantAuthorization, true, ""), nil, nil
	}
}

// SimulateMsgRevokeAuthorization generates a MsgRevokeAuthorization with random values
// nolint: interfacer
func SimulateMsgRevokeAuthorization(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		return simtypes.NoOpMsg(types.ModuleName), nil, nil
	}
}
