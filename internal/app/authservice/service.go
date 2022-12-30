package authservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/jwt"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type Service struct {
	api.UnimplementedAuthServiceServer
	repoGroup *repository.Group
	jwt       jwt.JWT
	crypto    crypto.Crypto
}

func NewService(
	repoGroup *repository.Group,
	jwt jwt.JWT,
	crypto crypto.Crypto,
) *Service {
	return &Service{
		repoGroup: repoGroup,
		jwt:       jwt,
		crypto:    crypto,
	}
}
