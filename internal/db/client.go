package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ClientImpl struct {
	db *sqlx.DB
}

func NewClient(cfg *config.Config) (*ClientImpl, error) {
	db, err := sqlx.Connect("pgx", cfg.PgDSN)
	if err != nil {
		return nil, fmt.Errorf("can't connect to pg instance, %v", err)
	}

	db.SetMaxOpenConns(cfg.PgMaxOpenConn)
	db.SetMaxIdleConns(cfg.PgIdleConn)

	go func() {
		t := time.NewTicker(cfg.PgPingInterval)

		for range t.C {
			if err := db.Ping(); err != nil {
				log.Errorf("error ping %v", err)
			}
		}
	}()

	// closer.Add(db.Close)
	return &ClientImpl{
		db: db,
	}, nil
}

type txKey struct{}

func (db *ClientImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.getClient(ctx).ExecContext(ctx, query, args...)
}

func (db *ClientImpl) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return db.getClient(ctx).QueryxContext(ctx, query, args...)
}

func (db *ClientImpl) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return db.getClient(ctx).QueryRowxContext(ctx, query, args...)
}

func (db *ClientImpl) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.getClient(ctx).GetContext(ctx, dest, query, args...)
}

func (db *ClientImpl) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.getClient(ctx).SelectContext(ctx, dest, query, args...)
}

func (db *ClientImpl) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return db.db.BeginTxx(ctx, opts)
}

func (db *ClientImpl) getClient(ctx context.Context) Client {
	if v := ctx.Value(txKey{}); v != nil {
		if tx, ok := v.(*sqlx.Tx); ok {
			return &txMock{Tx: tx}
		}
	}
	return &dbMock{db.db}
}

type txMock struct{ *sqlx.Tx }

func (t *txMock) BeginTxx(context.Context, *sql.TxOptions) (*sqlx.Tx, error) {
	return nil, errors.New("can't begin transaction in transaction")
}

type dbMock struct{ *sqlx.DB }

func (t *dbMock) Commit() error {
	return errors.New("can't commit db")
}

func (t *dbMock) Rollback() error {
	return errors.New("can't rollback db")
}
