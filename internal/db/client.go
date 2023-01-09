package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/cliffordsimak-76-cards/gophkeeper/internal/config"
)

func NewClient(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.PgDSN)
	if err != nil {
		return nil, fmt.Errorf("can't connect to pg instance, %v", err)
	}

	db.SetMaxOpenConns(cfg.PgMaxOpenConn)
	db.SetMaxIdleConns(cfg.PgIdleConn)

	return db, nil
}

func Ping(
	ctx context.Context,
	db *sqlx.DB,
	cfg *config.Config,
) {
	ticker := time.NewTicker(cfg.PgPingInterval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := db.PingContext(ctx); err != nil {
					log.Fatalf("error ping %v", err)
				}
			}
		}
	}()
}
