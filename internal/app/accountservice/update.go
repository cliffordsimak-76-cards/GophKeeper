package accountservice

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdateAccount(
	ctx context.Context,
	req *api.UpdateAccountRequest,
) (*api.Account, error) {
	err := validateUpdateRequest(req)
	if err != nil {
		log.Printf("error validate update account request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.GetUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	account := adapters.UpdateAccountRequestFromPb(req, userID)
	account, err = s.repoGroup.AccountRepository.Update(ctx, account)
	if err != nil {
		log.Printf("error update account in db: %s", err)
		return nil, status.Error(codes.Internal, "error update account in db")
	}

	return adapters.AccountToPb(account), nil
}

func validateUpdateRequest(req *api.UpdateAccountRequest) error {
	return validation.Errors{
		"login":    validation.Validate(req.GetLogin(), validation.Required),
		"password": validation.Validate(req.GetPassword(), validation.Required),
	}.Filter()
}
