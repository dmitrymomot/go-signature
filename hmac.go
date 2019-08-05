package signature

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
)

type hmacFunc func(securedInput, key []byte) string

func calculateHmac(securedInput, key []byte) string {
	hasher := hmac.New(func() hash.Hash { return sha1.New() }, key)
	hasher.Write(securedInput)
	return hex.EncodeToString(hasher.Sum(nil))
}

func calculateHmac256(securedInput, key []byte) string {
	hasher := hmac.New(func() hash.Hash { return sha256.New() }, key)
	hasher.Write(securedInput)
	return hex.EncodeToString(hasher.Sum(nil))
}

func validateHmac(securedInput, sign, key []byte, fn hmacFunc) error {
	sum := fn(securedInput, key)
	if !hmac.Equal([]byte(sum), sign) {
		return errors.New("invalid signature")
	}
	return nil
}
