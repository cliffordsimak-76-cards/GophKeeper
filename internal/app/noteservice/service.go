package noteservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type service struct {
	api.UnimplementedNoteServiceServer
	repoGroup *repository.Group
	auth      auth.Client
}

// NewService creates a new note manager
func NewService(
	repoGroup *repository.Group,
	auth auth.Client,
) *service {
	return &service{
		repoGroup: repoGroup,
		auth:      auth,
	}
}
