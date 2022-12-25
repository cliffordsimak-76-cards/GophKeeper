package cardservice

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Create(
	ctx context.Context,
	req *api.CreateCardRequest,
) (*api.Card, error) {
	err := validateCreateRequest(req)
	if err != nil {
		log.Printf("error validate create card request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.GetUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	card := adapters.CreateCardRequestFromPb(req, userID)
	card, err = s.repoGroup.CardRepository.Create(ctx, card)
	if err != nil {
		log.Printf("error create card in db: %s", err)
		return nil, status.Error(codes.Internal, "error create card in db")
	}

	return adapters.CardToPb(card), nil
}

func validateCreateRequest(req *api.CreateCardRequest) error {
	return validation.Errors{
		"name":   validation.Validate(req.GetName(), validation.Required),
		"number": validation.Validate(req.GetNumber(), validation.Required),
		"holder": validation.Validate(req.GetHolder(), validation.Required),
		"expire": validation.Validate(req.GetExpire(), validation.Required),
		"cvc":    validation.Validate(req.GetCvc(), validation.Required),
	}.Filter()
}
