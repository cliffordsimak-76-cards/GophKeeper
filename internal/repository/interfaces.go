//go:generate rm -rf ./mock_gen.go
//go:generate mockgen -destination=./mock_gen.go -package=repository -source=interfaces.go
package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
)

type Group struct {
	UserRepository    UserRepository
	CardRepository    CardRepository
	AccountRepository AccountRepository
	NoteRepository    NoteRepository
}

func NewGroup(db *sqlx.DB) *Group {
	return &Group{
		UserRepository:    NewUserRepository(db),
		CardRepository:    NewCardRepository(db),
		AccountRepository: NewAccountRepository(db),
		NoteRepository:    NewNoteRepository(db),
	}
}

type UserRepository interface {
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) error
	Get(context.Context, string) (*model.User, error)
}

type CardRepository interface {
	Create(context.Context, *model.Card) (*model.Card, error)
	Update(context.Context, *model.Card) (*model.Card, error)
	Get(context.Context, string) (*model.Card, error)
	List(context.Context, *CardListFilter) ([]*model.Card, error)
}

type AccountRepository interface {
	Create(context.Context, *model.Account) (*model.Account, error)
	Update(context.Context, *model.Account) (*model.Account, error)
	Get(context.Context, string) (*model.Account, error)
	List(context.Context, *AccountListFilter) ([]*model.Account, error)
}

type NoteRepository interface {
	Create(context.Context, *model.Note) (*model.Note, error)
	Update(context.Context, *model.Note) (*model.Note, error)
	Get(context.Context, string) (*model.Note, error)
	List(context.Context, *NoteListFilter) ([]*model.Note, error)
}
