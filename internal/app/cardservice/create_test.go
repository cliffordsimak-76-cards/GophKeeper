package cardservice

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_Create(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateCardRequest{}

		_, err := te.service.CreateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.CreateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		card := adapters.CreateCardRequestFromPb(req, userID)
		te.cardRepoMock.EXPECT().Create(te.ctx, card).
			Return(nil, errAny)

		_, err := te.service.CreateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		card := &model.Card{ID: "id"}
		te.cardRepoMock.EXPECT().Create(te.ctx, adapters.CreateCardRequestFromPb(req, userID)).
			Return(card, nil)

		response, err := te.service.CreateCard(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.CardToPb(card), response)
	})
}
