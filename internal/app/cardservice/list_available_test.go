package cardservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_ListAvailableCards(t *testing.T) {
	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableCardsRequest{}

		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.ListAvailableCards(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableCardsRequest{}

		userID := "user-id"
		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.CardListFilterFromPb(req, userID)
		te.cardRepoMock.EXPECT().List(te.ctx, filter).
			Return(nil, errAny)

		_, err := te.service.ListAvailableCards(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableCardsRequest{}

		userID := "user-id"
		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.CardListFilterFromPb(req, userID)
		cards := []*model.Card{
			{ID: "id-1"},
			{ID: "id-2"},
		}
		te.cardRepoMock.EXPECT().List(te.ctx, filter).
			Return(cards, nil)

		response, err := te.service.ListAvailableCards(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.ListAvailableCardsToPb(cards), response)
	})
}
