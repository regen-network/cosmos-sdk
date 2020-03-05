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
type PrivKeyNistp256 struct {
	PrivKey [PrivKeyNistp256Size]byte
}

func (privKey PrivKeyNistp256) PubKey() crypto.PubKey {
	PrivKeyFromBytes(privKey)
	return privKey.PubKey()
}

func (privKey PrivKeyNistp256) Sign(rand io.Reader, err error) {

}

func (privKey PrivKeyNistp256) Bytes() []byte {
	return nil
}

func (privKey PrivKeyNistp256) Equal() {

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
	privatekey, err := ecdsa.GenerateKey(pubKeyCurve, rand) // this generates a public & private key pair

	if err != nil {
		fmt.Println(err)
		return [256]byte{}, err
	}

	x, _ := cdc.MarshalBinaryBare(&privatekey.D)

	return x, nil
}
