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

const notesTName = "notes"

type NoteRepositoryImpl struct {
	db db.Client
}

func NewNoteRepositoryImpl(dbClient db.Client) *NoteRepositoryImpl {
	return &NoteRepositoryImpl{
		db: dbClient,
	}
}

func (r *NoteRepositoryImpl) Create(
	ctx context.Context,
	note *model.Note,
) (*model.Note, error) {
	note.CreatedAt = time.Now()
	query, _, err := goqu.Insert(notesTName).Rows(
		note,
	).Returning("id").ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for create note: %w", err)
	}

	var id string
	if err = r.db.QueryRowxContext(ctx, query).Scan(&id); err != nil {
		return nil, fmt.Errorf("can't create note: %w", err)
	}

	note.ID = id

	return note, nil
}

func (r *NoteRepositoryImpl) Update(
	ctx context.Context,
	note *model.Note,
) (*model.Note, error) {
	note.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updateDataset := goqu.Update(notesTName).Set(note).Where(goqu.I("id").Eq(note.ID))
	query, _, err := updateDataset.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("can't build query for update note: %w", err)
	}

	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return nil, fmt.Errorf("can't update note %w", err)
	}

	return note, nil
}

func (r *NoteRepositoryImpl) Get(
	ctx context.Context,
	id string,
) (*model.Note, error) {
	query, _, err := goqu.From(notesTName).Where(
		goqu.I("id").Eq(id),
	).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query get note: %w", err)
	}

	note := &model.Note{}
	if err = r.db.GetContext(ctx, note, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrEntityNotFound
		}
		return nil, fmt.Errorf("unable to get note: %w", err)
	}

	return note, nil
}

type NoteListFilter struct {
	UserID string
}

func (filter *NoteListFilter) toDataset() *goqu.SelectDataset {
	selectDataset := goqu.From(notesTName)

	if filter.UserID != "" {
		selectDataset = selectDataset.Where(goqu.I("user_id").In(filter.UserID))
	}

	return selectDataset
}

func (r *NoteRepositoryImpl) List(
	ctx context.Context,
	filter *NoteListFilter,
) ([]*model.Note, error) {
	noteList := make([]*model.Note, 0)

	query, _, err := filter.toDataset().ToSQL()
	if err != nil {
		return nil, fmt.Errorf("unable to create query list notes: %w", err)
	}

	if err = r.db.SelectContext(ctx, &noteList, query); err != nil {
		return nil, fmt.Errorf("unable to list notes: %w", err)
	}

	return noteList, nil
}
