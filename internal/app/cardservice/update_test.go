package cardservice

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_Update(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateCardRequest{}

		_, err := te.service.UpdateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("extract user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.UpdateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("error encrypt card", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		userID := "user-id"
		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		te.cryptoClientMock.EXPECT().Encrypt(gomock.Any()).
			Return("", errAny)

		_, err := te.service.UpdateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		userID := "user-id"
		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		te.cryptoClientMock.EXPECT().Encrypt(gomock.Any()).
			Return(gomock.Any().String(), nil).
			AnyTimes()

		te.cardRepoMock.EXPECT().Update(te.ctx, gomock.Any()).
			Return(nil, errAny)

		_, err := te.service.UpdateCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}

		userID := "user-id"
		te.authClientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		te.cryptoClientMock.EXPECT().Encrypt(gomock.Any()).
			Return(gomock.Any().String(), nil).
			AnyTimes()

		card := &model.Card{ID: "id"}
		te.cardRepoMock.EXPECT().Update(te.ctx, gomock.Any()).
			Return(card, nil)

		response, err := te.service.UpdateCard(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.CardToPb(card), response)
	})
}
