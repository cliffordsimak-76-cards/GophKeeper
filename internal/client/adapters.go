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
		Id:     card.ID,
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
		ID:     card.Id,
		Name:   card.Name,
		Number: card.Number,
		Expire: card.Expire,
		CVC:    card.Cvc,
		Holder: card.Holder,
	}
}

func buildCreateAccountRequest(account *Account) *api.CreateAccountRequest {
	return &api.CreateAccountRequest{
		Name:     account.Name,
		Login:    account.Login,
		Password: account.Password,
	}
}

func buildUpdateAccountRequest(account *Account) *api.UpdateAccountRequest {
	return &api.UpdateAccountRequest{
		Id:       account.ID,
		Name:     account.Name,
		Login:    account.Login,
		Password: account.Password,
	}
}

func buildListAvailableAccountsRequest() *api.ListAvailableAccountsRequest {
	return &api.ListAvailableAccountsRequest{}
}

func buildGetAccountRequest(id string) *api.GetAccountRequest {
	return &api.GetAccountRequest{
		Id: id,
	}
}

func pbAccountToAccount(account *api.Account) *Account {
	return &Account{
		ID:       account.Id,
		Name:     account.Name,
		Login:    account.Login,
		Password: account.Password,
	}
}

func buildCreateNoteRequest(note *Note) *api.CreateNoteRequest {
	return &api.CreateNoteRequest{
		Name: note.Name,
		Text: note.Text,
	}
}

func buildUpdateNoteRequest(note *Note) *api.UpdateNoteRequest {
	return &api.UpdateNoteRequest{
		Id:   note.ID,
		Name: note.Name,
		Text: note.Text,
	}
}

func buildListAvailableNotesRequest() *api.ListAvailableNotesRequest {
	return &api.ListAvailableNotesRequest{}
}

func buildGetNoteRequest(id string) *api.GetNoteRequest {
	return &api.GetNoteRequest{
		Id: id,
	}
}

func pbNoteToNote(note *api.Note) *Note {
	return &Note{
		ID:   note.Id,
		Name: note.Name,
		Text: note.Text,
	}
}
