package models

import (
	"database/sql"
	"time"
)

type Status struct {
	Id_status   int32
	Name        string	`json:"name" validate:"required"`
	Description string	`json:"description"`
	Created_at  time.Time
	Updated_at  sql.NullTime
	Deleted_at  sql.NullTime
}
