package signature_test

import (
	"testing"
	"time"

	"github.com/dmitrymomot/go-signature"
	"github.com/stretchr/testify/require"
)

func TestSigner_String(t *testing.T) {
	s := signature.NewSigner[string]([]byte("signing-key"))
	require.NotNil(t, s)

	t.Run("success", func(t *testing.T) {
		testData := "test123"

		token, err := s.Sign(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: temporary", func(t *testing.T) {
		testData := "test123"

		token, err := s.SignTemporary(testData, time.Second*5)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})
}

func TestSigner_Struct(t *testing.T) {
	type example struct {
		ID   uint64
		Text string
	}

	s := signature.NewSigner[example]([]byte("signing-key"))
	require.NotNil(t, s)

	t.Run("success", func(t *testing.T) {
		testData := example{
			ID:   123,
			Text: "test123",
		}

		token, err := s.Sign(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: temporary", func(t *testing.T) {
		testData := example{
			ID:   123,
			Text: "test123",
		}

		token, err := s.SignTemporary(testData, time.Second*5)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})
}

func TestSigner256_String(t *testing.T) {
	s := signature.NewSigner256[string]([]byte("signing-key"))
	require.NotNil(t, s)

	t.Run("success", func(t *testing.T) {
		testData := "test123"

		token, err := s.Sign(testData)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})

	t.Run("success: temporary", func(t *testing.T) {
		testData := "test123"

		token, err := s.SignTemporary(testData, time.Second*5)
		require.NoError(t, err)
		require.NotEmpty(t, token)

		data, err := s.Parse(token)
		require.NoError(t, err)
		require.Equal(t, testData, data)
	})
}
