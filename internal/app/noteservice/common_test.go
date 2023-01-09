package noteservice

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

	noteRepoMock *repository.MockNoteRepository
	clientMock   *auth.MockClient

	service *Service
}

func newTestEnv(t *testing.T) *testEnv {
	ctrl := gomock.NewController(t)
	te := &testEnv{
		ctx:          context.Background(),
		noteRepoMock: repository.NewMockNoteRepository(ctrl),
		clientMock:   auth.NewMockClient(ctrl),
	}

	repoGroup := &repository.Group{
		NoteRepository: te.noteRepoMock,
	}
	te.service = NewService(
		repoGroup,
		te.clientMock,
	)
	return te
}
