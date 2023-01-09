package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func Test_CreateAccountRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.CreateAccountRequest{}
		userID := "user-id"
		data := CreateAccountRequestFromPb(req, userID)
		expectedData := &model.Account{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.CreateAccountRequest{
			Name:     "name",
			Login:    "login",
			Password: "password",
			Metadata: "metadata",
		}
		userID := "user-id"
		data := CreateAccountRequestFromPb(req, userID)
		expectedData := &model.Account{
			Name:     "name",
			UserID:   userID,
			Login:    "login",
			Password: "password",
			Metadata: "metadata",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_UpdateAccountRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.UpdateAccountRequest{}
		userID := "user-id"
		data := UpdateAccountRequestFromPb(req, userID)
		expectedData := &model.Account{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.UpdateAccountRequest{
			Id:       "id",
			Name:     "name",
			Login:    "login",
			Password: "password",
			Metadata: "metadata",
		}
		userID := "user-id"
		data := UpdateAccountRequestFromPb(req, userID)
		expectedData := &model.Account{
			ID:       "id",
			Name:     "name",
			UserID:   userID,
			Login:    "login",
			Password: "password",
			Metadata: "metadata",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_AccountListFilterFromPb(t *testing.T) {
	t.Run("all data filler", func(t *testing.T) {
		req := &api.ListAvailableAccountsRequest{}
		userID := "user-id"
		data := AccountListFilterFromPb(req, userID)
		expectedData := &repository.AccountListFilter{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
}
