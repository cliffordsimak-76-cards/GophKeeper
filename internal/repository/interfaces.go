//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=repository -source=interfaces.go
package repository

import (
	"context"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/db"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
)

type Group struct {
	CardRepository CardRepository
	UserRepository UserRepository
}

func NewGroup(client db.Client) *Group {
	return &Group{
		CardRepository: NewCardRepositoryImpl(client),
		UserRepository: NewUserRepositoryImpl(client),
	}
}

type CardRepository interface {
	Create(context.Context, *model.Card) (*model.Card, error)
	Update(context.Context, *model.Card) (*model.Card, error)
	Get(context.Context, string) (*model.Card, error)
	List(context.Context, *CardListFilter) ([]*model.Card, error)
}

type UserRepository interface {
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) error
	Get(context.Context, string) (*model.User, error)
}
