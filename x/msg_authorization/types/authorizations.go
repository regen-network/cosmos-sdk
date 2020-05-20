package types

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AuthorizationI interface {
	MsgType() string
	Accept(msg sdk.Msg, block abci.Header) (allow bool, updated AuthorizationI, delete bool)
}
