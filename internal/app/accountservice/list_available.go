package accountservice

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/accountservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func (s *service) ListAvailableAccounts(
	ctx context.Context,
	req *api.ListAvailableAccountsRequest,
) (*api.ListAvailableAccountsResponse, error) {
	userID, err := s.auth.ExtractUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	filter := adapters.AccountListFilterFromPb(req, userID)
	accounts, err := s.repoGroup.AccountRepository.List(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "error list available accounts in db")
	}

	return adapters.ListAvailableAccountsToPb(accounts), nil
}
