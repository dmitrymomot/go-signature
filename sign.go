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
func New(data interface{}) (string, error) {
	return newToken(data, 0, calculateHmac)
}

// New256 returns sha256-signed string without expiration time
func New256(data interface{}) (string, error) {
	return newToken(data, 0, calculateHmac256)
}

// NewTemporary returns sha1-signed string with expiration time
// ttl - token life time in seconds
func NewTemporary(data interface{}, ttl int64) (string, error) {
	return newToken(data, ttl, calculateHmac)
}

// New256Temporary returns sha256-signed string with expiration time
// ttl - token life time in seconds
func New256Temporary(data interface{}, ttl int64) (string, error) {
	return newToken(data, ttl, calculateHmac256)
}

// Parse token signed by sha1
func Parse(token string) (interface{}, error) {
	return parseToken(token, calculateHmac)
}

// Parse256 parses sha256-signed token
func Parse256(token string) (interface{}, error) {
	return parseToken(token, calculateHmac256)
}

type tokenClaims struct {
	Payload   interface{} `json:"p"`
	ExpiresAt int64       `json:"e,omitempty"`
}

func newToken(data interface{}, ttl int64, fn hmacFunc) (string, error) {
	claims := &tokenClaims{Payload: data}
	if ttl > 0 {
		claims.ExpiresAt = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}
	b, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", base64Encode(b), base64Encode([]byte(fn(b, []byte(signingKey))))), nil
}

func parseToken(token string, fn hmacFunc) (interface{}, error) {
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 2 {
		return nil, errors.New("invalid token")
	}
	payload, err := base64Decode(tokenParts[0])
	if err != nil {
		return nil, errors.New("could not decode payload")
	}
	signature, err := base64Decode(tokenParts[1])
	if err != nil {
		return nil, errors.New("could not decode signature")
	}

	if err := validateHmac(payload, signature, []byte(signingKey), fn); err != nil {
		return nil, err
	}

	t := new(tokenClaims)
	if err := json.Unmarshal(payload, t); err != nil {
		return nil, err
	}

	if t.ExpiresAt > 0 {
		if time.Now().After(time.Unix(t.ExpiresAt, 0)) {
			return nil, errors.New("token expired")
		}
	}

	return t.Payload, nil
}
