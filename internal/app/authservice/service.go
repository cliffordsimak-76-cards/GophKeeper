package authservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

// Service is a user manager
type Service struct {
	api.UnimplementedAuthServiceServer
	repoGroup *repository.Group
	jwt       jwt.Client
	crypto    crypto.Client
}

// NewService creates a new user manager
func NewService(
	repoGroup *repository.Group,
	jwt jwt.Client,
	crypto crypto.Client,
) *Service {
	return &Service{
		repoGroup: repoGroup,
		jwt:       jwt,
		crypto:    crypto,
	}
}
