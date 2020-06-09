package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers concrete types and interfaces on the given codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgGrantAuthorization{}, "cosmos-sdk/GrantAuthorization", nil)
	cdc.RegisterConcrete(MsgRevokeAuthorization{}, "cosmos-sdk/RevokeAuthorization", nil)
	cdc.RegisterConcrete(MsgExecAuthorized{}, "cosmos-sdk/ExecDelegated", nil)
	cdc.RegisterConcrete(SendAuthorization{}, "cosmos-sdk/SendAuthorization", nil)
	cdc.RegisterConcrete(AuthorizationGrant{}, "cosmos-sdk/AuthorizationGrant", nil)
	cdc.RegisterConcrete(GenericAuthorization{}, "cosmos-sdk/GenericAuthorization", nil)

	cdc.RegisterInterface((*AuthorizationI)(nil), nil)
}

// RegisterInterfaces will registers the msgs and interfaces of the msg_auth module
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGrantAuthorization{},
		&MsgRevokeAuthorization{},
		&MsgExecAuthorized{},
	)

	registry.RegisterInterface(
		"cosmos_sdk.msgauth.v1.Authorization",
		(*AuthorizationI)(nil),
		&SendAuthorization{},
		// &GenericAuthorization{},
	)
}

var (
	amino = codec.New()

	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
