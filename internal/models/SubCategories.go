package models

import (
	"database/sql"
	"time"
)

type SubCategory struct {
	IDSubCate   int32        `json:"id_sub_cate"`
	IDCate      int32        `json:"id_cate" validate:"required"`
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
