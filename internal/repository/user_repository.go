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

const usersTName = "users"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *model.User,
) (*model.User, error) {
	user.CreatedAt = time.Now()
	query, _, err := goqu.Insert(usersTName).Rows(
		user,
	).Returning("id").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for create user: %w", err)
	}

	var id string
	if err = r.db.QueryRowContext(ctx, query).Scan(&id); err != nil {
		return nil, fmt.Errorf("can't create user: %w", err)
	}

	user.ID = id

	return user, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	user *model.User,
) error {
	user.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updateDataset := goqu.Update(usersTName).Set(user).Where(goqu.I("id").Eq(user.ID))
	query, _, err := updateDataset.ToSQL()
	if err != nil {
		return fmt.Errorf("can't build query for update user: %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("can't update user %v", err)
	}

	return nil
}

func (r *userRepository) Get(
	ctx context.Context,
	username string,
) (*model.User, error) {
	query, _, err := goqu.From(usersTName).Where(
		goqu.I("username").Eq(username),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query get user: %w", err)
	}

	user := &model.User{}
	if err = r.db.GetContext(ctx, user, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}
		return nil, fmt.Errorf("unable to get user: %w", err)
	}

	return user, nil
}
