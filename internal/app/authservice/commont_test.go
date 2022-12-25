package authservice

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

	userRepoMock *repository.MockUserRepository
	jwtMock      *auth.MockJWT
	cryptoMock   *crypto.MockCrypto

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:          context.Background(),
		userRepoMock: repository.NewMockUserRepository(ctrl),
		jwtMock:      auth.NewMockJWT(ctrl),
		cryptoMock:   crypto.NewMockCrypto(ctrl),
	}

	repoGroup := &repository.Group{
		UserRepository: te.userRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.jwtMock,
		te.cryptoMock,
	)
	return te
}
