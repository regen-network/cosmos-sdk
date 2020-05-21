package types

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func (authorization SendAuthorization) MsgType() string {
	return bank.MsgSend{}.Type()
}

func (authorization SendAuthorization) Accept(msg sdk.Msg, block abci.Header) (allow bool, updated AuthorizationI, delete bool) {
	switch msg := msg.(type) {
	case bank.MsgSend:
		limitLeft, isNegative := authorization.SpendLimit.SafeSub(msg.Amount)
		if isNegative {
			return false, nil, false
		}
		if limitLeft.IsZero() {
			return true, nil, true
		}

		return true, SendAuthorization{SpendLimit: limitLeft}, false
	}
	return false, nil, false
}
