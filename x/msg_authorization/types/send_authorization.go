package types

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func (authorization SendAuthorization) MsgType() string {
	return bank.MsgSend{}.Type()
}

func (authorization SendAuthorization) Accept(msg sdk.Msg, block abci.Header) (allow bool, updated *codectypes.Any, delete bool) {
	switch msg := msg.(type) {
	case *bank.MsgSend:
		limitLeft, isNegative := authorization.SpendLimit.SafeSub(msg.Amount)
		if isNegative {
			return false, nil, false
		}
		if limitLeft.IsZero() {
			return true, nil, true
		}

		authorization, err := ConvertToAny(SendAuthorization{SpendLimit: limitLeft})
		if err != nil {
			return false, nil, false
		}

		return true, authorization, false
	}
	return false, nil, false
}

// ConvertToAny converts interface(types.AuthorizationI) to any
func ConvertToAny(authorization AuthorizationI) (*codectypes.Any, error) {
	msg, ok := authorization.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("can't proto marshal %T", msg)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	return any, nil
}
