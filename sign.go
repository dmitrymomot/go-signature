package signature

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// SigningKey option
var signingKey string

// SetSigningKey helper to set up global signing key
func SetSigningKey(sk string) {
	signingKey = sk
}

// New returns sha1-signed string without expiration time
func New[Payload any](data Payload) (string, error) {
	return newToken(data, 0, calculateHmac)
}

// New256 returns sha256-signed string without expiration time.
func New256[Payload any](data Payload) (string, error) {
	return newToken(data, 0, calculateHmac256)
}

// NewTemporary returns sha1-signed string with expiration time
// ttl - token life time in seconds
func NewTemporary[Payload any](data Payload, ttl time.Duration) (string, error) {
	return newToken(data, ttl, calculateHmac)
}

// New256Temporary returns sha256-signed string with expiration time
// ttl - token life time in seconds
func New256Temporary[Payload any](data Payload, ttl time.Duration) (string, error) {
	return newToken(data, ttl, calculateHmac256)
}

// Parse token signed by sha1
func Parse[Payload any](token string) (Payload, error) {
	return parseToken[Payload](token, calculateHmac)
}

// Parse256 parses sha256-signed token
func Parse256[Payload any](token string) (Payload, error) {
	return parseToken[Payload](token, calculateHmac256)
}

type tokenClaims[Payload any] struct {
	Payload   Payload `json:"p"`
	ExpiresAt int64   `json:"e,omitempty"`
}

func newToken[Payload any](data Payload, ttl time.Duration, fn hmacFunc) (string, error) {
	claims := &tokenClaims[Payload]{Payload: data}
	if ttl > 0 {
		claims.ExpiresAt = time.Now().Add(ttl).Unix()
	}
	b, err := json.Marshal(claims)
	if err != nil {
		return "", errors.Join(ErrFailedToMarshalTokenClaims, err)
	}
	signedStr, err := fn(b, []byte(signingKey))
	if err != nil {
		return "", errors.Join(ErrInvalidSignature, err)
	}
	return fmt.Sprintf("%s.%s", base64Encode(b), base64Encode([]byte(signedStr))), nil
}

func parseToken[Payload any](token string, fn hmacFunc) (p Payload, err error) {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 2 {
		return p, errors.Join(ErrInvalidToken, ErrInvalidTokenFormat)
	}
	payload, err := base64Decode(tokenParts[0])
	if err != nil {
		return p, errors.Join(ErrInvalidToken, err)
	}
	signature, err := base64Decode(tokenParts[1])
	if err != nil {
		return p, errors.Join(ErrInvalidToken, err)
	}

	if err := validateHmac(payload, signature, []byte(signingKey), fn); err != nil {
		return p, errors.Join(ErrInvalidToken, err)
	}

	t := new(tokenClaims[Payload])
	if err := json.Unmarshal(payload, t); err != nil {
		return p, errors.Join(ErrInvalidToken, err)
	}

	if t.ExpiresAt > 0 {
		if time.Now().After(time.Unix(t.ExpiresAt, 0)) {
			return p, errors.Join(ErrInvalidToken, ErrTokenExpired)
		}
	}

	return t.Payload, nil
}
