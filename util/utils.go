package utils

import (
	"crypto/hmac"
	"crypto/sha1"
)

// Calculate HMAC1 signature of stringToSign using the given key and return
// true or false if the signature matches.
//
// This uses hmac.Equal for constant time comparison
func VerifyHMACSignature(signature []byte, stringToSign []byte, key []byte) bool {
	h := hmac.New(sha1.New, key)
	h.Write(stringToSign)
	return hmac.Equal(signature, h.Sum(nil))
}
