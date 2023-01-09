package cardservice

import (
	"context"
	"errors"
	"log"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func (s *Service) GetCard(
	ctx context.Context,
	req *api.GetCardRequest,
) (*api.Card, error) {
	err := validateGetRequest(req)
	if err != nil {
		log.Printf("error validate get card request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	card, err := s.repoGroup.CardRepository.Get(ctx, req.GetId())
	if err != nil {
		log.Printf("error get card in db: %s", err)
		if errors.Is(err, repository.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, "card is not found")
		}
		return nil, status.Error(codes.Internal, "error get card in db")
	}

	card, err = s.decryptCard(card)
	if err != nil {
		log.Printf("error decrypt card: %s", err)
		return nil, status.Error(codes.Internal, "error decrypt card")
	}

	return adapters.CardToPb(card), nil
}

func validateGetRequest(req *api.GetCardRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.GetId(), validation.Required, is.UUIDv4),
	}.Filter()
}
