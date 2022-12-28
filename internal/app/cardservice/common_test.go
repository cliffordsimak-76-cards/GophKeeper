package cardservice

import (
	"context"
	"errors"
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	"github.com/golang/mock/gomock"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx context.Context

	cardRepoMock *repository.MockCardRepository
	authMock     *auth.MockAuth
	cryptoMock   *crypto.MockCrypto

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:          context.Background(),
		cardRepoMock: repository.NewMockCardRepository(ctrl),
		authMock:     auth.NewMockAuth(ctrl),
		cryptoMock:   crypto.NewMockCrypto(ctrl),
	}

	repoGroup := &repository.Group{
		CardRepository: te.cardRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.authMock,
		te.cryptoMock,
	)
	return te
}
