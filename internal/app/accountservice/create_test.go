package accountservice

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Create(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateAccountRequest{}

		_, err := te.service.CreateAccount(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateAccountRequest{
			Name:     "name",
			Login:    "login",
			Password: "password",
		}

		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.CreateAccount(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateAccountRequest{
			Name:     "name",
			Login:    "login",
			Password: "password",
		}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		account := adapters.CreateAccountRequestFromPb(req, userID)
		te.accountRepoMock.EXPECT().Create(te.ctx, account).
			Return(nil, errAny)

		_, err := te.service.CreateAccount(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateAccountRequest{
			Name:     "name",
			Login:    "login",
			Password: "password",
		}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		account := &model.Account{ID: "id"}
		te.accountRepoMock.EXPECT().Create(te.ctx, adapters.CreateAccountRequestFromPb(req, userID)).
			Return(account, nil)

		response, err := te.service.CreateAccount(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.AccountToPb(account), response)
	})
}
