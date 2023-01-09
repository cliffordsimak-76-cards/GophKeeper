package accountservice

import (
	"context"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

// CreateAccount creates a login-password pair
func (s *Service) CreateAccount(
	ctx context.Context,
	req *api.CreateAccountRequest,
) (*api.Account, error) {
	err := validateCreateRequest(req)
	if err != nil {
		log.Printf("error validate create account request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.ExtractUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	account := adapters.CreateAccountRequestFromPb(req, userID)
	account, err = s.repoGroup.AccountRepository.Create(ctx, account)
	if err != nil {
		log.Printf("error create account in db: %s", err)
		return nil, status.Error(codes.Internal, "error create account in db")
	}

	return adapters.AccountToPb(account), nil
}

func validateCreateRequest(req *api.CreateAccountRequest) error {
	return validation.Errors{
		"login":    validation.Validate(req.GetLogin(), validation.Required),
		"password": validation.Validate(req.GetPassword(), validation.Required),
	}.Filter()
}
