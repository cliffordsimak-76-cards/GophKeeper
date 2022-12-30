package model

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        string       `db:"id" goqu:"skipinsert,skipupdate"`
	UserID    string       `db:"user_id"`
	Name      string       `db:"name"`
	Text      string       `db:"text"`
	Metadata  string       `db:"metadata"`
	CreatedAt time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
