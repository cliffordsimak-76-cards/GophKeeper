package noteservice

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_ListAvailableNotes(t *testing.T) {
	t.Run("get user id error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableNotesRequest{}

		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return("", errAny)

		_, err := te.service.ListAvailableNotes(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableNotesRequest{}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.NoteListFilterFromPb(req, userID)
		te.noteRepoMock.EXPECT().List(te.ctx, filter).
			Return(nil, errAny)

		_, err := te.service.ListAvailableNotes(te.ctx, req)
		require.Error(t, err)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListAvailableNotesRequest{}

		userID := "user-id"
		te.authMock.EXPECT().ExtractUserIdFromContext(te.ctx).
			Return(userID, nil)

		filter := adapters.NoteListFilterFromPb(req, userID)
		notes := []*model.Note{
			{ID: "id-1"},
			{ID: "id-2"},
		}
		te.noteRepoMock.EXPECT().List(te.ctx, filter).
			Return(notes, nil)

		response, err := te.service.ListAvailableNotes(te.ctx, req)
		require.NoError(t, err)
		require.Equal(t, adapters.ListAvailableNotesToPb(notes), response)
	})
}
