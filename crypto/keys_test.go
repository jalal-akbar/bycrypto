package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivKey(t *testing.T) {
	privKey := GeneratePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privateKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), publicKeyLen)
}
