package accountservice

import (
	"context"
	"errors"
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	"github.com/golang/mock/gomock"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx context.Context

	accountRepoMock *repository.MockAccountRepository
	authMock        *auth.MockAuth

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:             context.Background(),
		accountRepoMock: repository.NewMockAccountRepository(ctrl),
		authMock:        auth.NewMockAuth(ctrl),
	}

	repoGroup := &repository.Group{
		AccountRepository: te.accountRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.authMock,
	)
	return te
}
