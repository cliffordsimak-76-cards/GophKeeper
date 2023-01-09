package noteservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type Service struct {
	api.UnimplementedNoteServiceServer
	repoGroup *repository.Group
	auth      auth.Client
}

func NewService(
	repoGroup *repository.Group,
	auth auth.Client,
) *Service {
	return &Service{
		repoGroup: repoGroup,
		auth:      auth,
	}
}
