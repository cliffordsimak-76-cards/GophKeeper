package model

import (
	"database/sql"
	"time"
)

type Card struct {
	ID        string       `db:"id" goqu:"skipinsert,skipupdate"`
	UserID    string       `db:"user_id"`
	Name      string       `db:"name"`
	Number    string       `db:"number"`
	Holder    string       `db:"holder"`
	Expire    string       `db:"expire"`
	CVC       string       `db:"cvc"`
	CreatedAt time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
