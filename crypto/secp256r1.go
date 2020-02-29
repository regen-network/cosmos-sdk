package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"io"

	"github.com/tendermint/tendermint/crypto"
)

var pubKeyCurve = elliptic.P256()

const PrivKeyNistp256Size = 256

// PrivKeyNistp256 implements crypto.PrivKey.
type PrivKeyNistp256 [PrivKeyNistp256Size]byte

func (privKey PrivKeyNistp256) PubKey() crypto.PubKey {
	return nil
}

// GenPrivKey generates a new sr25519 private key.
// It uses OS randomness in conjunction with the current global random seed
// in tendermint/libs/common to generate the private key.
func GenPrivKey() PrivKeyNistp256 {
	privateKey, err := genPrivKey(crypto.CReader())
	if err != nil {
		fmt.Println(err)
	}
	return privateKey
}

func genPrivKey(rand io.Reader) (PrivKeyNistp256, error) {
	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubKeyCurve, rand) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		return [256]byte{}, err
	}

	//var privKeyEd PrivKeyNistp256
	//copy(privKeyEd[:], privatekey.D)

	return nil, nil
}
