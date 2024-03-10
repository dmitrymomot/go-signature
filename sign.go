package signature

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// tokenClaims represents the claims of a token.
type tokenClaims[Payload any] struct {
	Payload   Payload `json:"p"`
	ExpiresAt int64   `json:"e,omitempty"`
}

// NewToken generates a new token with the provided signing key, payload data, time-to-live (ttl),
// and HMAC function. It returns the generated token as a string and an error if any.
// The payload can be of any type, specified by the `Payload` type parameter.
// If the ttl is greater than zero, the token will have an expiration time set to the current time
// plus the ttl duration.
// The `fn` parameter is the HMAC function used to sign the token.
// It takes the token claims as a JSON byte array and the signing key as input,
// and returns the signed string and an error if any.
// The generated token is a combination of the base64-encoded token claims and the base64-encoded
// signed string, separated by a dot ('.').
func NewToken[Payload any](signingKey []byte, data Payload, ttl time.Duration, fn hmacFunc) (string, error) {
	claims := &tokenClaims[Payload]{Payload: data}
	if ttl > 0 {
		claims.ExpiresAt = time.Now().Add(ttl).Unix()
	}
	b, err := json.Marshal(claims)
	if err != nil {
		return "", errors.Join(ErrFailedToMarshalTokenClaims, err)
	}
	signedStr, err := fn(b, signingKey)
	if err != nil {
		return "", errors.Join(ErrInvalidSignature, err)
	}
	return fmt.Sprintf("%s.%s", base64Encode(b), base64Encode([]byte(signedStr))), nil
}

// ParseToken parses a token and returns the payload contained within it.
// It takes a signing key, a token string, a HMAC function, and a validation function as input parameters.
// The signing key is used to verify the token's signature.
// The token string is expected to be in the format "payload.signature".
// The HMAC function is used to calculate the signature of the token.
// The validation function is used to validate the payload and signature against the signing key and HMAC function.
// The payload is decoded from base64 and unmarshaled into a tokenClaims struct.
// If the token is invalid or expired, an error is returned.
// Otherwise, the payload is returned.
func ParseToken[Payload any](signingKey []byte, token string, fn hmacFunc, vfn validateFunc) (p Payload, err error) {
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

	if err := vfn(payload, signature, []byte(signingKey), fn); err != nil {
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
