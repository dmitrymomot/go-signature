package signature_test

import (
	"testing"
	"time"

	"github.com/dmitrymomot/go-signature"
	"github.com/stretchr/testify/require"
)

func TestSha1(t *testing.T) {
	signingKey := "signing-key"
	signature.SetSigningKey(signingKey)

	t.Run("success: string", func(t *testing.T) {
		testData := "test123"

		token, err := signature.New(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse[string](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: int", func(t *testing.T) {
		testData := 123

		token, err := signature.New(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse[int](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: struct", func(t *testing.T) {
		type example struct {
			ID   uint64
			Text string
		}
		testData := example{ID: 123, Text: "test123"}

		token, err := signature.New(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse[example](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
		require.Equal(t, testData.ID, data.ID)
	})

	t.Run("invalid signature", func(t *testing.T) {
		testData := "test123"

		token, err := signature.New(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		signature.SetSigningKey("invalid-key")
		data, err := signature.Parse[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("invalid token format", func(t *testing.T) {
		token := "invalid-token"
		_, err := signature.Parse[string](token)
		require.ErrorIs(t, err, signature.ErrInvalidTokenFormat)
	})

	t.Run("invalid base64 string", func(t *testing.T) {
		_, err := signature.Parse[string]("invalid==.token")
		require.ErrorIs(t, err, signature.ErrInvalidToken)

		_, err = signature.Parse[string]("invalid.token++@")
		require.ErrorIs(t, err, signature.ErrInvalidToken)
	})
}

func TestSha256(t *testing.T) {
	signingKey := "signing-key"
	signature.SetSigningKey(signingKey)

	t.Run("success", func(t *testing.T) {
		testData := "test123"

		token, err := signature.New256(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse256[string](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		testData := "test123"

		token, err := signature.New256(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		signature.SetSigningKey("invalid-key")
		data, err := signature.Parse256[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})
}

func TestSha1Temp(t *testing.T) {
	signingKey := "signing-key"
	signature.SetSigningKey(signingKey)

	t.Run("success", func(t *testing.T) {
		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewTemporary(testData, ttl)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse[string](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.NewTemporary(testData, ttl)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		signature.SetSigningKey("invalid-key")
		data, err := signature.Parse[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("expired token", func(t *testing.T) {
		testData := "test123"

		token, err := signature.NewTemporary(testData, 1)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrTokenExpired)
	})
}

func TestSha256Temp(t *testing.T) {
	signingKey := "signing-key"
	signature.SetSigningKey(signingKey)

	t.Run("success", func(t *testing.T) {
		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.New256Temporary(testData, ttl)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse256[string](token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("invalid signature", func(t *testing.T) {
		testData := "test123"
		ttl := time.Second * 5

		token, err := signature.New256Temporary(testData, ttl)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		signature.SetSigningKey("invalid-key")
		data, err := signature.Parse256[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrInvalidSignature)
	})

	t.Run("expired token", func(t *testing.T) {
		testData := "test123"

		token, err := signature.New256Temporary(testData, 1)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := signature.Parse256[string](token)
		require.Error(t, err)
		require.Empty(t, data)
		require.ErrorIs(t, err, signature.ErrTokenExpired)
	})
}
