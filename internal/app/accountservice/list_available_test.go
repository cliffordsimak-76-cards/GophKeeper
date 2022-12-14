package accountservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_ListAvailableAccounts(t *testing.T) {
	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableAccountsRequest{}

		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.ListAvailableAccounts(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableAccountsRequest{}

		userID := "user-id"
		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.AccountListFilterFromPb(req, userID)
		te.accountRepoMock.EXPECT().List(te.ctx, filter).
			Return(nil, errAny)

		_, err := te.service.ListAvailableAccounts(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableAccountsRequest{}

		userID := "user-id"
		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.AccountListFilterFromPb(req, userID)
		accounts := []*model.Account{
			{ID: "id-1"},
			{ID: "id-2"},
		}
		te.accountRepoMock.EXPECT().List(te.ctx, filter).
			Return(accounts, nil)

		response, err := te.service.ListAvailableAccounts(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.ListAvailableAccountsToPb(accounts), response)
	})
}
