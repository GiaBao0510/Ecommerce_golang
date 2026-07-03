package models

type Role struct {
	Role_id     int32
	Role_name   string `json:"role_name" validate:"required"`
	Description string `json:"description"`
}
