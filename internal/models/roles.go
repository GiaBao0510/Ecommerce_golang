package models

type Role struct {
	Role_id     int32
	Role_name   string `json:"role_name" validate:"required"`
	Description string `json:"description"`
}


type CreateRoleRequest struct {
	Role_name   string `json:"role_name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdateRolePutRequest struct {
	Role_name   string `json:"role_name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=500"`
}

type UpdateRolePatchRequest struct {
	Role_name   *string `json:"role_name" binding:"omitempty,required,min=2,max=100"`
	Description *string `json:"description" binding:"omitempty,max=500"`
}