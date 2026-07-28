package models

import (
	"database/sql"
	"time"
)

type Status struct {
	Id_status   int32
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Created_at  time.Time
	Updated_at  sql.NullTime
	Deleted_at  sql.NullTime
}

type CreateStatusRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdateStatusPutRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdateStatusPatchRequest struct {
	Name        *string `json:"name" binding:"omitempty,required,min=2,max=100"`
	Description *string `json:"description" binding:"omitempty,max=500"`
}
