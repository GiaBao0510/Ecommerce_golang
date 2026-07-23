package models

import (
	"database/sql"
	"time"
)

type Role_Permission struct {
	Action_id     int32 `json:"action_id"`
	Role_id       int32 `json:"role_id" validate:"required"`
	Permission_id int32 `json:"permission_id" validate:"required"`
	Created_at    time.Time
	Updated_at    sql.NullTime
	Deleted_at    sql.NullTime
}
