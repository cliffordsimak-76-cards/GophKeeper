package adapters

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func CreateAccountRequestFromPb(req *api.CreateAccountRequest, userID string) *model.Account {
	return &model.Account{
		Name:     req.GetName(),
		UserID:   userID,
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
		Metadata: req.GetMetadata(),
	}
}

func UpdateAccountRequestFromPb(req *api.UpdateAccountRequest, userID string) *model.Account {
	return &model.Account{
		ID:       req.GetId(),
		Name:     req.GetName(),
		UserID:   userID,
		Login:    req.GetLogin(),
		Password: req.GetPassword(),
		Metadata: req.GetMetadata(),
	}
}

func AccountListFilterFromPb(_ *api.ListAvailableAccountsRequest, userID string) *repository.AccountListFilter {
	return &repository.AccountListFilter{
		UserID: userID,
	}
}
