package noteservice

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListAvailableNotes(
	ctx context.Context,
	req *api.ListAvailableNotesRequest,
) (*api.ListAvailableNotesResponse, error) {
	userID, err := s.auth.ExtractUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	filter := adapters.NoteListFilterFromPb(req, userID)
	notes, err := s.repoGroup.NoteRepository.List(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "error list available notes in db")
	}

	return adapters.ListAvailableNotesToPb(notes), nil
}
