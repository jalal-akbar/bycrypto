package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privateKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), publicKeyLen)
}

func TestPrivateKeyToString(t *testing.T) {
	addressStr := "2dd842f22584657887a52ea56b35173b979a0cce"
	seed := "49fb9069e2f7932a758636ecae29f36831a5936a60ebe91ccfaeb672ab857420"
	privKey := NewPrivateKeyToString(seed)
	assert.Equal(t, privateKeyLen, len(privKey.Bytes()))

	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
	//fmt.Println(address)
}

func TestPrivKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("jalal foe bar")

	sign := privKey.Sign(msg)
	assert.True(t, sign.Verify(pubKey, msg), msg)
	fmt.Println(pubKey)

	// test with invalid msg
	assert.False(t, sign.Verify(pubKey, []byte("jalal")))
	// test with invalid pubkey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()

	assert.False(t, sign.Verify(invalidPubKey, msg))
	fmt.Println("invalid", invalidPubKey)

}

func TestPublicKeyToAddres(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	assert.Equal(t, addressLen, len(address.Bytes()))
}
