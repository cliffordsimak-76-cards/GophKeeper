package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
)

func Test_AccountToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Account{}
		mapedData := AccountToPb(data)
		expectedData := &api.Account{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Account{
			ID:        "id",
			Name:      "name",
			Login:     "login",
			Password:  "password",
		
		}
		mapedData := AccountToPb(data)
		expectedData := &api.Account{
			Id:       "id",
			Name:     "name",
			Login:   "login",
			Password: "password",
		}
		require.Equal(t, expectedData, mapedData)
	})
}

func Test_ListAvailableAccountsToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := []*model.Account{}
		mapedData := ListAvailableAccountsToPb(data)
		expectedData := &api.ListAvailableAccountsResponse{
			Accounts: make([]*api.AvailableAccount, 0),
		}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := []*model.Account{
			{
				ID:   "id-1",
				Name: "name-1",
			},
			{
				ID:   "id-2",
				Name: "name-2",
			},
		}
		mapedData := ListAvailableAccountsToPb(data)
		expectedData := &api.ListAvailableAccountsResponse{
			Accounts: []*api.AvailableAccount{
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

func Test_AvailableAccountToPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		data := &model.Account{}
		mapedData := AvailableAccountToPb(data)
		expectedData := &api.AvailableAccount{}
		require.Equal(t, expectedData, mapedData)
	})
	t.Run("all data filler", func(t *testing.T) {
		data := &model.Account{
			ID:   "id",
			Name: "name",
		}
		mapedData := AvailableAccountToPb(data)
		expectedData := &api.AvailableAccount{
			Id:   "id",
			Name: "name",
		}
		require.Equal(t, expectedData, mapedData)
	})
}
