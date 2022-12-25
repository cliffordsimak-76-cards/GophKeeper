package cardservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type Service struct {
	api.UnimplementedCardServiceServer
	repoGroup *repository.Group
	auth      auth.Auth
}

func NewService(
	repoGroup *repository.Group,
	auth auth.Auth,
) *Service {
	return &Service{
		repoGroup: repoGroup,
		auth:      auth,
	}
}
