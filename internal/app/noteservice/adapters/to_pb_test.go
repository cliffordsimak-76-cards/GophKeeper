package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
)

func Test_NoteToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Note{}
		mapedData := NoteToPb(data)
		expectedData := &api.Note{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Note{
			ID:       "id",
			Name:     "name",
			Text:     "text",
			Metadata: "metadata",
		}
		mapedData := NoteToPb(data)
		expectedData := &api.Note{
			Id:       "id",
			Name:     "name",
			Text:     "text",
			Metadata: "metadata",
		}
		require.Equal(t, expectedData, mapedData)
	})
}

func Test_ListAvailableNotesToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := []*model.Note{}
		mapedData := ListAvailableNotesToPb(data)
		expectedData := &api.ListAvailableNotesResponse{
			Notes: make([]*api.AvailableNote, 0),
		}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := []*model.Note{
			{
				ID:   "id-1",
				Name: "name-1",
			},
			{
				ID:   "id-2",
				Name: "name-2",
			},
		}
		mapedData := ListAvailableNotesToPb(data)
		expectedData := &api.ListAvailableNotesResponse{
			Notes: []*api.AvailableNote{
				{
					Id:   "id-1",
					Name: "name-1",
				},
				{
					Id:   "id-2",
					Name: "name-2",
				},
			},
		}
		require.Equal(t, expectedData, mapedData)
	})
}

func Test_AvailableNoteToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Note{}
		mapedData := AvailableNoteToPb(data)
		expectedData := &api.AvailableNote{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Note{
			ID:   "id",
			Name: "name",
		}
		mapedData := AvailableNoteToPb(data)
		expectedData := &api.AvailableNote{
			Id:   "id",
			Name: "name",
		}
		require.Equal(t, expectedData, mapedData)
	})
}
