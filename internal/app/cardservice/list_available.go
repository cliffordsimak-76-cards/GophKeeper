package cardservice

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/cardservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func (s *Service) ListAvailableCards(
	ctx context.Context,
	req *api.ListAvailableCardsRequest,
) (*api.ListAvailableCardsResponse, error) {
	userID, err := s.auth.ExtractUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	filter := adapters.CardListFilterFromPb(req, userID)
	cards, err := s.repoGroup.CardRepository.List(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "error list available cards in db")
	}

	return adapters.ListAvailableCardsToPb(cards), nil
}
