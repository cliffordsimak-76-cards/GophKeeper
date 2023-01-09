package jwt

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

type testEnv struct {
	ctx       context.Context
	secretKey string
	jwt       *client
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{
		ctx:       context.Background(),
		secretKey: "secret",
	}

	te.jwt = &client{
		secretKey:     te.secretKey,
		tokenDuration: 1 * time.Minute,
	}
	return te
}

func Test_NewJWTImpl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cfg := &config.Config{}
		require.NotNil(t, NewClient(cfg))
	})
}

func Test_Verify(t *testing.T) {
	t.Run("error unexpected signing method", func(t *testing.T) {
		te := newTestEnv(t)

		ps256token := "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.e30.U-" +
			"Cb0WOtN22xWM8LXX26Y21Fq-YWiYWhBwwwT9aekKMHEowAjkdOfjo4CQqOf" +
			"famP-ihDBKAwLJA5RPTAUAU5nuNCqStCkKJMH-cLAVen1t9Ls2VCJQCaIiv" +
			"PcmW8yoFJg5hkTfCBGvVFN979K5p2MJQt0HSBN3wpyRZjZb8ULc3MgZAonU" +
			"ANG2zT_hDrVRGSRY4HrGBC2ux8NGk8O4vk6S168Ihvy3QRBxlwhNYlG_Zu1" +
			"uWHsq-gagzkcYS9JuIRi9Ftf6MKEOwCo7Poh14POlzJpTWONO4t79MNjzCh" +
			"1zx8zHOT07A5Y9VFY9YG8Xn3lc9mvdTsDWZGErjxw8Nng"
		err := te.jwt.Verify(ps256token)

		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error unexpected signing method"))
	})
	t.Run("error parse jwt token", func(t *testing.T) {
		te := newTestEnv(t)

		wrongSecretToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9" +
			"HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U"
		err := te.jwt.Verify(wrongSecretToken)

		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error parse jwt token"))
	})
	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJoZW" +
			"xsbyJ9.GqJuYcGhFOL5hgQLWM3XrKvQg2MektSq3vKTJPtK_sA"
		err := te.jwt.Verify(validToken)

		require.NoError(t, err)
	})
}

func Test_ExtractUserID(t *testing.T) {
	t.Run("error unexpected signing method", func(t *testing.T) {
		te := newTestEnv(t)

		ps256token := "eyJhbGciOiJQUzI1NiIsInR5cCI6IkpXVCJ9.e30.U-" +
			"Cb0WOtN22xWM8LXX26Y21Fq-YWiYWhBwwwT9aekKMHEowAjkdOfjo4CQqOf" +
			"famP-ihDBKAwLJA5RPTAUAU5nuNCqStCkKJMH-cLAVen1t9Ls2VCJQCaIiv" +
			"PcmW8yoFJg5hkTfCBGvVFN979K5p2MJQt0HSBN3wpyRZjZb8ULc3MgZAonU" +
			"ANG2zT_hDrVRGSRY4HrGBC2ux8NGk8O4vk6S168Ihvy3QRBxlwhNYlG_Zu1" +
			"uWHsq-gagzkcYS9JuIRi9Ftf6MKEOwCo7Poh14POlzJpTWONO4t79MNjzCh" +
			"1zx8zHOT07A5Y9VFY9YG8Xn3lc9mvdTsDWZGErjxw8Nng"
		_, err := te.jwt.ExtractUserID(ps256token)

		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error unexpected signing method"))
	})
	t.Run("error parse jwt token", func(t *testing.T) {
		te := newTestEnv(t)

		wrongSecretToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.Et9" +
			"HFtf9R3GEMA0IICOfFMVXY7kkTX1wr4qCyhIf58U"
		_, err := te.jwt.ExtractUserID(wrongSecretToken)

		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error parse jwt token"))
	})
	t.Run("error no userID in token", func(t *testing.T) {
		te := newTestEnv(t)

		noUserIDclaim := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemA" +
			"Ct8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"
		_, err := te.jwt.ExtractUserID(noUserIDclaim)

		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error no userID in token"))
	})
	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJoZW" +
			"xsbyJ9.GqJuYcGhFOL5hgQLWM3XrKvQg2MektSq3vKTJPtK_sA"
		userID, err := te.jwt.ExtractUserID(validToken)

		expectedUserID := "hello"
		require.NoError(t, err)
		require.Equal(t, expectedUserID, userID)
	})
}
