package noteservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_Create(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateNoteRequest{}

		_, err := te.service.CreateNote(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateNoteRequest{
			Name: "name",
			Text: "text",
		}

		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.CreateNote(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateNoteRequest{
			Name: "name",
			Text: "text",
		}

		userID := "user-id"
		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		note := adapters.CreateNoteRequestFromPb(req, userID)
		te.noteRepoMock.EXPECT().Create(te.ctx, note).
			Return(nil, errAny)

		_, err := te.service.CreateNote(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateNoteRequest{
			Name: "name",
			Text: "text",
		}

		userID := "user-id"
		te.clientMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		note := &model.Note{ID: "id"}
		te.noteRepoMock.EXPECT().Create(te.ctx, adapters.CreateNoteRequestFromPb(req, userID)).
			Return(note, nil)

		response, err := te.service.CreateNote(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.NoteToPb(note), response)
	})
}
