package models

import (
	"database/sql"
	"time"
)

type Permission struct {
	Action_id   int32
	Action_name string `json:"action_name" validate:"required"`
	Description string `json:"description"`
	Created_at  time.Time
	Updated_at  sql.NullTime
	Deleted_at  sql.NullTime
}