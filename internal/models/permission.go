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

type CreatePermissionRequest struct {
	Action_name string `json:"action_name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdatePermissionPutRequest struct {
	Action_name string `json:"action_name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdatePermissionPatchRequest struct {
	Action_name *string `json:"action_name" binding:"omitempty,required,min=2,max=100"`
	Description *string `json:"description" binding:"omitempty,max=500"`
}