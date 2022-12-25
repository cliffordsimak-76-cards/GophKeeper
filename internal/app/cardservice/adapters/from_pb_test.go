package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/stretchr/testify/require"
)

func Test_CreateCardRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.CreateCardRequest{}
		userID := "user-id"
		data := CreateCardRequestFromPb(req, userID)
		expectedData := &model.Card{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.CreateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}
		userID := "user-id"
		data := CreateCardRequestFromPb(req, userID)
		expectedData := &model.Card{
			Name:   "name",
			UserID: userID,
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			CVC:    "cvc",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_UpdateCardRequestFromPb(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := &api.UpdateCardRequest{}
		userID := "user-id"
		data := UpdateCardRequestFromPb(req, userID)
		expectedData := &model.Card{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
	t.Run("all data filler", func(t *testing.T) {
		req := &api.UpdateCardRequest{
			Name:   "name",
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			Cvc:    "cvc",
		}
		userID := "user-id"
		data := UpdateCardRequestFromPb(req, userID)
		expectedData := &model.Card{
			Name:   "name",
			UserID: userID,
			Number: "number",
			Holder: "holder",
			Expire: "expire",
			CVC:    "cvc",
		}
		require.Equal(t, expectedData, data)
	})
}

func Test_CardListFilterFromPb(t *testing.T) {
	t.Run("all data filler", func(t *testing.T) {
		req := &api.ListAvailableCardsRequest{}
		userID := "user-id"
		data := CardListFilterFromPb(req, userID)
		expectedData := &repository.CardListFilter{
			UserID: userID,
		}
		require.Equal(t, expectedData, data)
	})
}
