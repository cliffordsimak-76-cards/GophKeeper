package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
)

const cardsTName = "cards"

type cardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) *cardRepository {
	return &cardRepository{db}
}

func (r *cardRepository) Create(
	ctx context.Context,
	card *model.Card,
) (*model.Card, error) {
	card.CreatedAt = time.Now()
	query, _, err := goqu.Insert(cardsTName).Rows(
		card,
	).Returning("id").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for create card: %w", err)
	}

	var id string
	if err = r.db.QueryRowContext(ctx, query).Scan(&id); err != nil {
		return nil, fmt.Errorf("can't create card: %w", err)
	}

	card.ID = id

	return card, nil
}

func (r *cardRepository) Update(
	ctx context.Context,
	card *model.Card,
) (*model.Card, error) {
	card.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updateDataset := goqu.Update(cardsTName).Set(card).Where(goqu.I("id").Eq(card.ID))
	query, _, err := updateDataset.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for update card: %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return nil, fmt.Errorf("can't update card %w", err)
	}

	return card, nil
}

func (r *cardRepository) Get(
	ctx context.Context,
	id string,
) (*model.Card, error) {
	query, _, err := goqu.From(cardsTName).Where(
		goqu.I("id").Eq(id),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query get card: %w", err)
	}

	card := &model.Card{}
	if err = r.db.GetContext(ctx, card, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}
		return nil, fmt.Errorf("unable to get card: %w", err)
	}

	return card, nil
}

type CardListFilter struct {
	UserID string
}

func (filter *CardListFilter) toDataset() *goqu.SelectDataset {
	selectDataset := goqu.From(cardsTName)

	if filter.UserID != "" {
		selectDataset = selectDataset.Where(goqu.I("user_id").In(filter.UserID))
	}

	return selectDataset
}

func (r *cardRepository) List(
	ctx context.Context,
	filter *CardListFilter,
) ([]*model.Card, error) {
	cardList := make([]*model.Card, 0)

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query list cards: %w", err)
	}

	if err = r.db.SelectContext(ctx, &cardList, query); err != nil {
		return nil, fmt.Errorf("unable to list cards: %w", err)
	}

	return cardList, nil
}
