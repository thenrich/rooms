package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"encoding/hex"
)

func TestVerifyHMAC(t *testing.T) {
	key := []byte("thisismyhmackey")
	stringToSign := []byte("thisismystringtosign")
	hexSignature, _ := hex.DecodeString("332fe2cf10b0b9bb677ca51b74a6d14882917ab5")

	assert.True(t, VerifyHMACSignature(hexSignature, stringToSign, key))
	assert.False(t, VerifyHMACSignature([]byte("badsignature"), stringToSign, key))

}