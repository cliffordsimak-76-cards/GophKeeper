package client

import api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"

func buildRegisterRequest(username string, password string) *api.RegisterRequest {
	return &api.RegisterRequest{
		Username: username,
		Password: password,
	}
}

func buildLoginRequest(username string, password string) *api.LoginRequest {
	return &api.LoginRequest{
		Username: username,
		Password: password,
	}
}

func buildCreateCardRequest(card *Card) *api.CreateCardRequest {
	return &api.CreateCardRequest{
		Name:   card.Name,
		Number: card.Number,
		Expire: card.Expire,
		Cvc:    card.CVC,
		Holder: card.Holder,
	}
}

func buildUpdateCardRequest(card *Card) *api.UpdateCardRequest {
	return &api.UpdateCardRequest{
		Name:   card.Name,
		Number: card.Number,
		Expire: card.Expire,
		Cvc:    card.CVC,
		Holder: card.Holder,
	}
}

func buildListAvailableCardsRequest() *api.ListAvailableCardsRequest {
	return &api.ListAvailableCardsRequest{}
}

func buildGetCardRequest(id string) *api.GetCardRequest {
	return &api.GetCardRequest{
		Id: id,
	}
}

func pbCardToCard(card *api.Card) *Card {
	return &Card{
		Name:   card.Name,
		Number: card.Number,
		Expire: card.Expire,
		CVC:    card.Cvc,
		Holder: card.Holder,
	}
}
