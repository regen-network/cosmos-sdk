package secp256r1

import (
	"crypto/ecdsa"
	"io"

	"github.com/cosmos/cosmos-sdk/crypto"
)

func NewPrivateKey(rand io.Reader) {
	privatekey, err := ecdsa.GenerateKey(pubKeyCurve, rand) // this generates a public & private key pair

}
