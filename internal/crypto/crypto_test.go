package crypto

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

type testEnv struct {
	crypto *crypto
}

func newTestEnv(t *testing.T) *testEnv {
	cfg := &config.Config{
		EncryptPassword: "LrjOs4X9dmUFswtxmbsw9hKs2xqgAwxL",
	}

	return &testEnv{
		crypto: NewClient(cfg),
	}
}

func Test_IsCorrectPassword(t *testing.T) {
	t.Run("wrong password", func(t *testing.T) {
		te := newTestEnv(t)

		hash, _ := bcrypt.GenerateFromPassword([]byte("some password"), bcrypt.MinCost)

		require.False(t, te.crypto.IsCorrectPassword(string(hash), "other password"))
	})
	t.Run("correct password", func(t *testing.T) {
		te := newTestEnv(t)

		pwd := "some password"
		hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

		require.True(t, te.crypto.IsCorrectPassword(string(hash), pwd))
	})
}

func Test_Encrypt_Decrypt(t *testing.T) {
	t.Run("success encrypt decrypt", func(t *testing.T) {
		te := newTestEnv(t)

		str := "some string to encrypt"
		encrypted, err := te.crypto.Encrypt(str)
		require.NoError(t, err)

		decrypted, err := te.crypto.Decrypt(encrypted)
		require.NoError(t, err)

		require.Equal(t, decrypted, str)
	})
}
