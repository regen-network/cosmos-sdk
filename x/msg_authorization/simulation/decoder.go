package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/msg_authorization/types"
	tmkv "github.com/tendermint/tendermint/libs/kv"
)

// MsgAuthUnmarshaler sld
type MsgAuthUnmarshaler interface {
	UnmarshalAuthorization(bz []byte) (types.AuthorizationI, error)
}

// NewDecodeStore returns a function closure that unmarshals the KVPair's values
// to the corresponding types.
func NewDecodeStore(cdc MsgAuthUnmarshaler) func(kvA, kvB tmkv.Pair) string {
	return func(kvA, kvB tmkv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], []byte{0x00}):
			msgAuthA, err := cdc.UnmarshalAuthorization(kvA.Value)
			if err != nil {
				panic(err)
			}

			msgAuthB, err := cdc.UnmarshalAuthorization(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", msgAuthA, msgAuthB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
