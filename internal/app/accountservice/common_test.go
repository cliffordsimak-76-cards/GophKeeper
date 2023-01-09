package accountservice

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
)

var errAny = errors.New("any error")

type testEnv struct {
	ctx context.Context

	accountRepoMock *repository.MockAccountRepository
	clientMock      *auth.MockClient

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:             context.Background(),
		accountRepoMock: repository.NewMockAccountRepository(ctrl),
		clientMock:      auth.NewMockClient(ctrl),
	}

	repoGroup := &repository.Group{
		AccountRepository: te.accountRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.clientMock,
	)
	return te
}
