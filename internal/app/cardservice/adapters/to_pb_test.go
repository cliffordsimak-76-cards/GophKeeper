package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
)

func Test_CardToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Card{}
		mapedData := CardToPb(data)
		expectedData := &api.Card{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Card{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			CVC:    "cvc",
		}
		mapedData := CardToPb(data)
		expectedData := &api.Card{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}
		require.Equal(t, expectedData, mapedData)
	})
}

func Test_ListAvailableCardsToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := []*model.Card{}
		mapedData := ListAvailableCardsToPb(data)
		expectedData := &api.ListAvailableCardsResponse{
			Cards: make([]*api.AvailableCard, 0),
		}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := []*model.Card{
			{
				ID:   "id-1",
				Name: "name-1",
			},
			{
				ID:   "id-2",
				Name: "name-2",
			},
		}
		mapedData := ListAvailableCardsToPb(data)
		expectedData := &api.ListAvailableCardsResponse{
			Cards: []*api.AvailableCard{
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

func Test_AvailableCardToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Card{}
		mapedData := AvailableCardToPb(data)
		expectedData := &api.AvailableCard{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Card{
			ID:   "id",
			Name: "name",
		}
		mapedData := AvailableCardToPb(data)
		expectedData := &api.AvailableCard{
			Id:   "id",
			Name: "name",
		}
		require.Equal(t, expectedData, mapedData)
	})
}
