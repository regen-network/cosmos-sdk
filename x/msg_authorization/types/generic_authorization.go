package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (cap GenericAuthorization) MsgType() string {
	var msg sdk.Msg
	ModuleCdc.UnpackAny(&cap.Message, &msg)
	return msg.Type()
}

func (cap GenericAuthorization) Accept(msg sdk.Msg, block abci.Header) (allow bool, updated *codectypes.Any, delete bool) {
	genAuth, err := ConvertToAny(cap)
	if err != nil {
		return false, nil, false
	}
	return true, genAuth, false
}
