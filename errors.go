package signature

import "errors"

// Predefined errors
var (
	ErrInvalidSignature           = errors.New("invalid signature")
	ErrorCalculatingHmac          = errors.New("error calculating hmac")
	ErrFailedToMarshalTokenClaims = errors.New("failed to marshal token claims")
	ErrInvalidToken               = errors.New("invalid token")
	ErrTokenExpired               = errors.New("token expired")
	ErrInvalidTokenFormat         = errors.New("invalid token format")
)
