package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func NoteToPb(note *model.Note) *api.Note {
	return &api.Note{
		Id:       note.ID,
		Name:     note.Name,
		Text:     note.Text,
		Metadata: note.Metadata,
	}
}

func ListAvailableNotesToPb(notes []*model.Note) *api.ListAvailableNotesResponse {
	aNotes := make([]*api.AvailableNote, 0, len(notes))
	for _, c := range notes {
		aNotes = append(aNotes, AvailableNoteToPb(c))
	}
	return &api.ListAvailableNotesResponse{
		Notes: aNotes,
	}
}

func AvailableNoteToPb(note *model.Note) *api.AvailableNote {
	return &api.AvailableNote{
		Id:   note.ID,
		Name: note.Name,
	}
}
