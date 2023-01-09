package model

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        string       `db:"id" goqu:"skipinsert,skipupdate"`
	UserID    string       `db:"user_id"`
	Name      string       `db:"name"`
	Login     string       `db:"login"`
	Password  string       `db:"password"`
	Metadata  string       `db:"metadata"`
	CreatedAt time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
