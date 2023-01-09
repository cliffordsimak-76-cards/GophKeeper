package noteservice

import (
	"context"
	"errors"
	"log"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice/adapters"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

func (s *Service) GetNote(
	ctx context.Context,
	req *api.GetNoteRequest,
) (*api.Note, error) {
	err := validateGetRequest(req)
	if err != nil {
		log.Printf("error validate get note request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	note, err := s.repoGroup.NoteRepository.Get(ctx, req.GetId())
	if err != nil {
		log.Printf("error get note in db: %s", err)
		if errors.Is(err, repository.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, "note is not found")
		}
		return nil, status.Error(codes.Internal, "error get note in db")
	}

	return adapters.NoteToPb(note), nil
}

func validateGetRequest(req *api.GetNoteRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.GetId(), validation.Required, is.UUIDv4),
	}.Filter()
}
