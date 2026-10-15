package models

import (
	"database/sql"
)

type Brand struct {
	IDBrand   int32        `json:"id_brand"`
	Name      string       `json:"name" validate:"required"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
