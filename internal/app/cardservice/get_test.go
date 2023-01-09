package cardservice

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_Get(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetCardRequest{}

		_, err := te.service.GetCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetCardRequest{
			Id: "461ab5dc-9400-4a8b-9231-f98887406708",
		}

		te.cardRepoMock.EXPECT().Get(te.ctx, req.GetId()).
			Return(nil, errAny)

		_, err := te.service.GetCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("not found error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetCardRequest{
			Id: "461ab5dc-9400-4a8b-9231-f98887406708",
		}

		te.cardRepoMock.EXPECT().Get(te.ctx, req.GetId()).
			Return(nil, repository.ErrEntityNotFound)

		_, err := te.service.GetCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("error decrypt card", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetCardRequest{
			Id: "461ab5dc-9400-4a8b-9231-f98887406708",
		}

		card := &model.Card{ID: "id"}
		te.cardRepoMock.EXPECT().Get(te.ctx, req.GetId()).
			Return(card, nil)

		te.cryptoClientMock.EXPECT().Decrypt(gomock.Any()).
			Return("", errAny)

		_, err := te.service.GetCard(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetCardRequest{
			Id: "461ab5dc-9400-4a8b-9231-f98887406708",
		}

		card := &model.Card{ID: "id"}
		te.cardRepoMock.EXPECT().Get(te.ctx, req.GetId()).
			Return(card, nil)

		te.cryptoClientMock.EXPECT().Decrypt(gomock.Any()).
			Return(gomock.Any().String(), nil).
			AnyTimes()

		response, err := te.service.GetCard(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.CardToPb(card), response)
	})
}
