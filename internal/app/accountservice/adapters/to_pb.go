package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func AccountToPb(account *model.Account) *api.Account {
	return &api.Account{
		Id:       account.ID,
		Name:     account.Name,
		Login:    account.Login,
		Password: account.Password,
		Metadata: account.Metadata,
	}
}

func ListAvailableAccountsToPb(accounts []*model.Account) *api.ListAvailableAccountsResponse {
	aAccounts := make([]*api.AvailableAccount, 0, len(accounts))
	for _, c := range accounts {
		aAccounts = append(aAccounts, AvailableAccountToPb(c))
	}
	return &api.ListAvailableAccountsResponse{
		Accounts: aAccounts,
	}
}

func AvailableAccountToPb(account *model.Account) *api.AvailableAccount {
	return &api.AvailableAccount{
		Id:   account.ID,
		Name: account.Name,
	}
}
