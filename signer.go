package signature

import (
	"time"
)

// Signer is an interface that defines the methods for signing and parsing data.
type Signer[Payload any] interface {
	// Sign generates a signature for the given data.
	// It returns the generated signature as a string and any error encountered.
	Sign(data Payload) (string, error)

	// SignTemporary generates a temporary signature for the given data with a specified time-to-live (TTL).
	// It returns the generated signature as a string and any error encountered.
	SignTemporary(data Payload, ttl time.Duration) (string, error)

	// Parse parses the given token and returns the payload associated with it.
	// It returns the parsed payload and any error encountered.
	Parse(token string) (Payload, error)
}

// signer represents a signer that can generate and validate signatures for a given payload.
type signer[Payload any] struct {
	signingKey   []byte       // The key used for signing the payload.
	signFunc     hmacFunc     // The signing function used to generate the signature.
	validateFunc validateFunc // The validation function used to verify the signature.
}

// NewSigner creates a new instance of the Signer type with the specified signing key.
// The signing key is used to generate and validate HMAC signatures.
// The generic type parameter `Payload` represents the type of the payload that will be signed.
// The function returns a pointer to the created Signer instance.
func NewSigner[Payload any](signingKey []byte) Signer[Payload] {
	return &signer[Payload]{
		signingKey:   signingKey,
		signFunc:     CalculateHmac,
		validateFunc: ValidateHmac,
	}
}

// NewSigner256 creates a new instance of the Signer interface that uses HMAC-SHA256 for signing.
// It takes a signingKey as input, which is the secret key used for generating the HMAC signature.
// The generic type parameter `Payload` represents the type of the payload that will be signed.
// The function returns a pointer to the Signer implementation.
func NewSigner256[Payload any](signingKey []byte) Signer[Payload] {
	return &signer[Payload]{
		signingKey:   signingKey,
		signFunc:     CalculateHmac256,
		validateFunc: ValidateHmac,
	}
}

func (s *signer[Payload]) Sign(data Payload) (string, error) {
	return NewToken(s.signingKey, data, 0, s.signFunc)
}

func (s *signer[Payload]) SignTemporary(data Payload, ttl time.Duration) (string, error) {
	return NewToken(s.signingKey, data, ttl, s.signFunc)
}

func (s *signer[Payload]) Parse(token string) (Payload, error) {
	return ParseToken[Payload](s.signingKey, token, s.signFunc, s.validateFunc)
}
