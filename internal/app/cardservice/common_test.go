package cardservice

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx context.Context

	cardRepoMock     *repository.MockCardRepository
	authClientMock   *auth.MockClient
	cryptoClientMock *crypto.MockClient

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:              context.Background(),
		cardRepoMock:     repository.NewMockCardRepository(ctrl),
		authClientMock:   auth.NewMockClient(ctrl),
		cryptoClientMock: crypto.NewMockClient(ctrl),
	}

	repoGroup := &repository.Group{
		CardRepository: te.cardRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.authClientMock,
		te.cryptoClientMock,
	)
	return te
}
