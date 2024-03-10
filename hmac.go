package signature

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"hash"
)

type (
	hmacFunc     func(securedInput, key []byte) (string, error)
	validateFunc func(securedInput, sign, key []byte, fn hmacFunc) error
)

func CalculateHmac(securedInput, key []byte) (string, error) {
	hasher := hmac.New(func() hash.Hash { return sha1.New() }, key)
	if _, err := hasher.Write(securedInput); err != nil {
		return "", errors.Join(ErrorCalculatingHmac, err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func CalculateHmac256(securedInput, key []byte) (string, error) {
	hasher := hmac.New(func() hash.Hash { return sha256.New() }, key)
	if _, err := hasher.Write(securedInput); err != nil {
		return "", errors.Join(ErrorCalculatingHmac, err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func ValidateHmac(securedInput, sign, key []byte, fn hmacFunc) error {
	sum, err := fn(securedInput, key)
	if err != nil {
		return errors.Join(ErrInvalidSignature, err)
	}
	if !hmac.Equal([]byte(sum), sign) {
		return ErrInvalidSignature
	}
	return nil
}
