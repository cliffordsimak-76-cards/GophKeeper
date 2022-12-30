package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func CreateNoteRequestFromPb(req *api.CreateNoteRequest, userID string) *model.Note {
	return &model.Note{
		Name:     req.GetName(),
		UserID:   userID,
		Text:     req.GetText(),
		Metadata: req.GetMetadata(),
	}
}

func UpdateNoteRequestFromPb(req *api.UpdateNoteRequest, userID string) *model.Note {
	return &model.Note{
		ID:       req.GetId(),
		Name:     req.GetName(),
		UserID:   userID,
		Text:     req.GetText(),
		Metadata: req.GetMetadata(),
	}
}

func NoteListFilterFromPb(_ *api.ListAvailableNotesRequest, userID string) *repository.NoteListFilter {
	return &repository.NoteListFilter{
		UserID: userID,
	}
}
