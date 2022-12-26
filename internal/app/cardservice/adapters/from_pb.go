package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func CreateCardRequestFromPb(req *api.CreateCardRequest, userID string) *model.Card {
	return &model.Card{
		Name:   req.GetName(),
		UserID: userID,
		Number: req.GetNumber(),
		Holder: req.GetHolder(),
		Expire: req.GetExpire(),
		CVC:    req.GetCvc(),
	}
}

func UpdateCardRequestFromPb(req *api.UpdateCardRequest, userID string) *model.Card {
	return &model.Card{
		ID:     req.GetId(),
		Name:   req.GetName(),
		UserID: userID,
		Number: req.GetNumber(),
		Holder: req.GetHolder(),
		Expire: req.GetExpire(),
		CVC:    req.GetCvc(),
	}
}

func CardListFilterFromPb(_ *api.ListAvailableCardsRequest, userID string) *repository.CardListFilter {
	return &repository.CardListFilter{
		UserID: userID,
	}
}
