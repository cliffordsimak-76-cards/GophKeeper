//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=auth -source=interfaces.go
package auth

import (
	"context"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
)

type Auth interface {
	GetUserIdFromContext(ctx context.Context) (string, error)
}

type JWT interface {
	Generate(user *model.User) (string, error)
	Verify(accessToken string) error
}
