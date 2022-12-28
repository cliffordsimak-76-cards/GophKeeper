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

func (s *Service) UpdateNote(
	ctx context.Context,
	req *api.UpdateNoteRequest,
) (*api.Note, error) {
	err := validateUpdateRequest(req)
	if err != nil {
		log.Printf("error validate update note request: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.auth.GetUserIdFromContext(ctx)
	if err != nil {
		log.Printf("error get userID from context: %s", err)
		return nil, status.Error(codes.Internal, "error get userID from context")
	}

	note := adapters.UpdateNoteRequestFromPb(req, userID)
	note, err = s.repoGroup.NoteRepository.Update(ctx, note)
	if err != nil {
		log.Printf("error update note in db: %s", err)
		return nil, status.Error(codes.Internal, "error update note in db")
	}

	return adapters.NoteToPb(note), nil
}

func validateUpdateRequest(req *api.UpdateNoteRequest) error {
	return validation.Errors{
		"text": validation.Validate(req.GetText(), validation.Required),
	}.Filter()
}
