package cardservice

import (
	"fmt"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/auth"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/crypto"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/repository"
	api "github.com/cliffordsimak-76-cards/gophkeeper/pkg/gophkeeper-api"
)

type service struct {
	api.UnimplementedCardServiceServer
	repoGroup *repository.Group
	auth      auth.Client
	crypto    crypto.Client
}

// NewService creates a new card service
func NewService(
	repoGroup *repository.Group,
	auth auth.Client,
	crypto crypto.Client,
) *service {
	return &service{
		repoGroup: repoGroup,
		auth:      auth,
		crypto:    crypto,
	}
}

func (s *service) encryptCard(card *model.Card) (*model.Card, error) {
	encryptedCardNumber, err := s.crypto.Encrypt(card.Number)
	if err != nil {
		return nil, fmt.Errorf("error encrypt card number: %w", err)
	}

	encryptedCVC, err := s.crypto.Encrypt(card.CVC)
	if err != nil {
		return nil, fmt.Errorf("error encrypt card CVC: %w", err)
	}

	card.Number = encryptedCardNumber
	card.CVC = encryptedCVC
	return card, nil
}

func (s *service) decryptCard(card *model.Card) (*model.Card, error) {
	decryptedCardNumber, err := s.crypto.Decrypt(card.Number)
	if err != nil {
		return nil, fmt.Errorf("error decrypt card number: %w", err)
	}

	decryptedCVC, err := s.crypto.Decrypt(card.CVC)
	if err != nil {
		return nil, fmt.Errorf("error decrypt card CVC: %w", err)
	}

	card.Number = decryptedCardNumber
	card.CVC = decryptedCVC
	return card, nil
}
