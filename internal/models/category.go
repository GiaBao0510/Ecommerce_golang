package models

import (
	"database/sql"
)

type Category struct {
	IDCate      int32        `json:"id_cate"`
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"description"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
