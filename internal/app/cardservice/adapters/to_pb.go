package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func CardToPb(card *model.Card) *api.Card {
	return &api.Card{
		Id:     card.ID,
		Name:   card.Name,
		Number: card.Number,
		Holder: card.Holder,
		Expire: card.Expire,
		Cvc:    card.CVC,
	}
}

func ListAvailableCardsToPb(cards []*model.Card) *api.ListAvailableCardsResponse {
	aCards := make([]*api.AvailableCard, 0, len(cards))
	for _, c := range cards {
		aCards = append(aCards, AvailableCardToPb(c))
	}
	return &api.ListAvailableCardsResponse{
		Cards: aCards,
	}
}

func AvailableCardToPb(card *model.Card) *api.AvailableCard {
	return &api.AvailableCard{
		Id:   card.ID,
		Name: card.Name,
	}
}
