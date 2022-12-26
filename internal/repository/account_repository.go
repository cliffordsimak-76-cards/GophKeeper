package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/db"
	"github.com/cliffordsimak-76-cards/gophkeeper/internal/model"
	"github.com/doug-martin/goqu/v9"
)

const accountsTName = "accounts"

type AccountRepositoryImpl struct {
	db db.Client
}

func NewAccountRepositoryImpl(dbClient db.Client) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		db: dbClient,
	}
}

func (r *AccountRepositoryImpl) Create(
	ctx context.Context,
	account *model.Account,
) (*model.Account, error) {
	account.CreatedAt = time.Now()
	query, _, err := goqu.Insert(accountsTName).Rows(
		account,
	).Returning("id").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for create account: %w", err)
	}

	var id string
	if err = r.db.QueryRowxContext(ctx, query).Scan(&id); err != nil {
		return nil, fmt.Errorf("can't create account: %w", err)
	}

	account.ID = id

	return account, nil
}

func (r *AccountRepositoryImpl) Update(
	ctx context.Context,
	account *model.Account,
) (*model.Account, error) {
	account.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updateDataset := goqu.Update(accountsTName).Set(account).Where(goqu.I("id").Eq(account.ID))
	query, _, err := updateDataset.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for update account: %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return nil, fmt.Errorf("can't update account %w", err)
	}

	return account, nil
}

func (r *AccountRepositoryImpl) Get(
	ctx context.Context,
	id string,
) (*model.Account, error) {
	query, _, err := goqu.From(accountsTName).Where(
		goqu.I("id").Eq(id),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query get account: %w", err)
	}

	account := &model.Account{}
	if err = r.db.GetContext(ctx, account, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}
		return nil, fmt.Errorf("unable to get account: %w", err)
	}

	return account, nil
}

type AccountListFilter struct {
	UserID string
}

func (filter *AccountListFilter) toDataset() *goqu.SelectDataset {
	selectDataset := goqu.From(accountsTName)

	if filter.UserID != "" {
		selectDataset = selectDataset.Where(goqu.I("user_id").In(filter.UserID))
	}

	return selectDataset
}

func (r *AccountRepositoryImpl) List(
	ctx context.Context,
	filter *AccountListFilter,
) ([]*model.Account, error) {
	accountList := make([]*model.Account, 0)

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query list accounts: %w", err)
	}

	if err = r.db.SelectContext(ctx, &accountList, query); err != nil {
		return nil, fmt.Errorf("unable to list accounts: %w", err)
	}

	return accountList, nil
}
