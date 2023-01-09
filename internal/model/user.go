package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID             string       `db:"id" goqu:"skipinsert,skipupdate"`
	Username       string       `db:"username"`
	HashedPassword string       `db:"hasedpassword"`
	CreatedAt      time.Time    `db:"created_at" goqu:"skipupdate"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
}
