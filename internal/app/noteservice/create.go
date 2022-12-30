package noteservice

import (
	"context"
	"log"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/app/noteservice/adapters"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateNote(
	ctx context.Context,
	req *api.CreateNoteRequest,
) (*api.Note, error) {
	err := validateCreateRequest(req)
	if err != nil {
		log.Printf("error validate create note request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.ExtractUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	note := adapters.CreateNoteRequestFromPb(req, userID)
	note, err = s.repoGroup.NoteRepository.Create(ctx, note)
	if err != nil {
		log.Printf("error create note in db: %s", err)
		return nil, status.Error(codes.Internal, "error create note in db")
	}

	return adapters.NoteToPb(note), nil
}

func validateCreateRequest(req *api.CreateNoteRequest) error {
	return validation.Errors{
		"text": validation.Validate(req.GetText(), validation.Required),
	}.Filter()
}
