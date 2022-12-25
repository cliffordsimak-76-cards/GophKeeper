package authservice

import (
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type Service struct {
	api.UnimplementedAuthServiceServer
	repoGroup *repository.Group
	jwt       auth.JWT
	crypto    crypto.Crypto
}

func NewService(
	repoGroup *repository.Group,
	jwt auth.JWT,
	crypto    crypto.Crypto,
) *Service {
	return &Service{
		repoGroup: repoGroup,
		jwt:       jwt,
		crypto: crypto,
	}
}
