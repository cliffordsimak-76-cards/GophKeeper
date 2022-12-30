package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
)

func Test_CreateNoteRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.CreateNoteRequest{}
		userID := "user-id"
		data := CreateNoteRequestFromPb(req, userID)
		expectedData := &model.Note{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.CreateNoteRequest{
			Name:     "name",
			Text:     "text",
			Metadata: "metadata",
		}
		userID := "user-id"
		data := CreateNoteRequestFromPb(req, userID)
		expectedData := &model.Note{
			Name:     "name",
			UserID:   userID,
			Text:     "text",
			Metadata: "metadata",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_UpdateNoteRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.UpdateNoteRequest{}
		userID := "user-id"
		data := UpdateNoteRequestFromPb(req, userID)
		expectedData := &model.Note{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.UpdateNoteRequest{
			Id:       "id",
			Name:     "name",
			Text:     "text",
			Metadata: "metadata",
		}
		userID := "user-id"
		data := UpdateNoteRequestFromPb(req, userID)
		expectedData := &model.Note{
			ID:       "id",
			Name:     "name",
			UserID:   userID,
			Text:     "text",
			Metadata: "metadata",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_NoteListFilterFromPb(t *testing.T) {
	t.Run("all data filler", func(t *testing.T) {
		req := &api.ListAvailableNotesRequest{}
		userID := "user-id"
		data := NoteListFilterFromPb(req, userID)
		expectedData := &repository.NoteListFilter{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
}
