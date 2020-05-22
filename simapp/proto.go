package simapp

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	registry = types.NewInterfaceRegistry()

	protoCdc = codec.NewProtoCodec(registry)
)
