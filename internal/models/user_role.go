package models

type UserRole struct {
	Id_role int32  `json:"id_role" binding:"required"`
	Uuid    string `json:"uuid" binding:"required,uuid"`
}

type UserByRole struct {
	Uuid      string `json:"uid"`
	User_name string `json:"user_name"`
	Email     string `json:"email"`
	Phone_num string `json:"phone_num"`
	Address   string `json:"address"`
}

type RoleByUser struct {
	Role_id     int32  `json:"role_id"`
	Role_name   string `json:"role_name"`
	Description string `json:"description"`
}
