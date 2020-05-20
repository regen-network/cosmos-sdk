package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type AuthorizationI interface {
	MsgType() string
	Accept(msg sdk.Msg, block abci.Header) (allow bool, updated AuthorizationI, delete bool)
}
