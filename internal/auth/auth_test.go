package auth

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx         context.Context
	clientMock  *jwt.MockClient
	interceptor *client
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)

	te := &testEnv{
		ctx:        context.Background(),
		clientMock: jwt.NewMockClient(ctrl),
	}

	te.interceptor = NewClient(te.clientMock)
	return te
}

func Test_authorize(t *testing.T) {
	t.Run("error authorize", func(t *testing.T) {
		te := newTestEnv(t)

		err := te.interceptor.authorize(te.ctx)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error authorize"))
	})
	t.Run("access token is invalid", func(t *testing.T) {
		te := newTestEnv(t)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"

		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs(authHeader, token),
		)

		te.clientMock.EXPECT().Verify(token).
			Return(errAny)

		err := te.interceptor.authorize(te.ctx)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "access token is invalid"))
	})
	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"

		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs(authHeader, token),
		)

		te.clientMock.EXPECT().Verify(token).
			Return(nil)

		err := te.interceptor.authorize(te.ctx)
		require.NoError(t, err)
	})
}

func Test_ExtractUserIdFromContext(t *testing.T) {
	t.Run("error get token from context", func(t *testing.T) {
		te := newTestEnv(t)

		_, err := te.interceptor.ExtractUserIdFromContext(te.ctx)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error get token from context"))
	})
	t.Run("error extract userID from token", func(t *testing.T) {
		te := newTestEnv(t)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"

		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs(authHeader, token),
		)

		te.clientMock.EXPECT().ExtractUserID(token).
			Return("", errAny)

		_, err := te.interceptor.ExtractUserIdFromContext(te.ctx)
		require.Error(t, err)
		require.True(t, strings.Contains(err.Error(), "error extract userID from token"))
	})
	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"

		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs(authHeader, token),
		)

		userID := "userID"
		te.clientMock.EXPECT().ExtractUserID(token).
			Return(userID, nil)

		result, err := te.interceptor.ExtractUserIdFromContext(te.ctx)
		require.NoError(t, err)
		require.Equal(t, userID, result)
	})
}

func Test_getTokenFromContext(t *testing.T) {
	t.Run("error metadata is not provided", func(t *testing.T) {
		te := newTestEnv(t)

		_, err := extractTokenFromContext(te.ctx)

		require.Error(t, err)
		require.Equal(t, codes.Unauthenticated, status.Code(err))
	})
	t.Run("error authorization token is not provided", func(t *testing.T) {
		te := newTestEnv(t)
		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs("some key", "some value"),
		)

		_, err := extractTokenFromContext(te.ctx)

		require.Error(t, err)
		require.Equal(t, codes.Unauthenticated, status.Code(err))
	})
	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M"

		te.ctx = metadata.NewIncomingContext(
			te.ctx,
			metadata.Pairs(authHeader, testToken),
		)

		token, err := extractTokenFromContext(te.ctx)
		require.NoError(t, err)
		require.Equal(t, testToken, token)
	})
}
