package accountservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type service struct {
	api.UnimplementedAccountServiceServer
	repoGroup *repository.Group
	auth      auth.Client
}

// NewService creates a new login-password manager
func NewService(
	repoGroup *repository.Group,
	auth auth.Client,
) *service {
	return &service{
		repoGroup: repoGroup,
		auth:      auth,
	}
}
