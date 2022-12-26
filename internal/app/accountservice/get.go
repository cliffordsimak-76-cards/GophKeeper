package accountservice

import (
	"context"
	"errors"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Get(
	ctx context.Context,
	req *api.GetAccountRequest,
) (*api.Account, error) {
	err := validateGetRequest(req)
	if err != nil {
		log.Printf("error validate get account request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	account, err := s.repoGroup.AccountRepository.Get(ctx, req.GetId())
	if err != nil {
		log.Printf("error get account in db: %s", err)
		if errors.Is(err, repository.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, "account is not found")
		}
		return nil, status.Error(codes.Internal, "error get account in db")
	}

	return adapters.AccountToPb(account), nil
}

func validateGetRequest(req *api.GetAccountRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.GetId(), validation.Required, is.UUIDv4),
	}.Filter()
}
