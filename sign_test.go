package signature_test

import (
	"testing"
	"time"

	"github.com/dmitrymomot/go-signature"
	"github.com/stretchr/testify/require"
)

func TestSha1(t *testing.T) {
	t.Parallel()

	signingKey := []byte("signing-key")

	t.Run("success: string", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: int", func(t *testing.T) {
		t.Parallel()

		testData := 123

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[int](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: struct", func(t *testing.T) {
		t.Parallel()

		type example struct {
			ID   uint64
			Text string
		}
		testData := example{ID: 123, Text: "test123"}

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[example](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
		require.Equal(t, testData.ID, data.ID)
	})

	t.Run("invalid signature", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string]([]byte("invalid-key"), token, signature.CalculateHmac, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("invalid token format", func(t *testing.T) {
		t.Parallel()

		token := "invalid-token"
		_, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.ErrorIs(t, err, signature.ErrInvalidTokenFormat)
	})

	t.Run("invalid base64 string", func(t *testing.T) {
		t.Parallel()

		token := "invalid==.token"
		_, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.ErrorIs(t, err, signature.ErrInvalidToken)

		token = "invalid.token++@"
		_, err = signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.ErrorIs(t, err, signature.ErrInvalidToken)
	})
}

func TestSha256(t *testing.T) {
	t.Parallel()

	signingKey := []byte("signing-key")

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac256)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac256, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 0, signature.CalculateHmac256)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string]([]byte("invalid-key"), token, signature.CalculateHmac256, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})
}

func TestSha1Temp(t *testing.T) {
	t.Parallel()

	signingKey := []byte("signing-key")

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewToken(signingKey, testData, ttl, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		t.Parallel()

		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewToken(signingKey, testData, ttl, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string]([]byte("invalid-key"), token, signature.CalculateHmac, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("expired token", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 1, signature.CalculateHmac)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrTokenExpired)
	})
}

func TestSha256Temp(t *testing.T) {
	t.Parallel()

	signingKey := []byte("signing-key")

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewToken(signingKey, testData, ttl, signature.CalculateHmac256)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac256, signature.ValidateHmac)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		t.Parallel()

		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewToken(signingKey, testData, ttl, signature.CalculateHmac256)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string]([]byte("invalid-key"), token, signature.CalculateHmac256, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("expired token", func(t *testing.T) {
		t.Parallel()

		testData := "test123"

		token, err := signature.NewToken(signingKey, testData, 1, signature.CalculateHmac256)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.ParseToken[string](signingKey, token, signature.CalculateHmac256, signature.ValidateHmac)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrTokenExpired)
	})
}
