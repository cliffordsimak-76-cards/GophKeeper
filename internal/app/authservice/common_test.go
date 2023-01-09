package authservice

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx context.Context

	userRepoMock     *repository.MockUserRepository
	jwtClientMock    *jwt.MockClient
	cryptoClientMock *crypto.MockClient

	service *service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:              context.Background(),
		userRepoMock:     repository.NewMockUserRepository(ctrl),
		jwtClientMock:    jwt.NewMockClient(ctrl),
		cryptoClientMock: crypto.NewMockClient(ctrl),
	}

	repoGroup := &repository.Group{
		UserRepository: te.userRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.jwtClientMock,
		te.cryptoClientMock,
	)
	return te
}
