package models

import (
	"database/sql"
	"time"
)

type Users struct {
	Uuid              string `json:"uuid"`
	Id_status         int32  `json:"id_status"`
	User_name         string `json:"user_name"`
	Birth_date        time.Time
	Email             string `json:"email"`
	Phone_num         string `json:"phone_num"`
	Is_email_verified bool   `json:"is_email_verified"`
	Is_phone_verified bool   `json:"is_phone_verified"`
	Address           string `json:"address"`
	Password_hash     string `json:"password_hash"`
	Avatar_url        string `json:"avatar_url"`
	Created_at        time.Time
	Updated_at        sql.NullTime
	Deleted_at        sql.NullTime
}

type CreateUsersRequest struct {
	Id_status     int32     `json:"id_status" binding:"required"`
	User_name     string    `json:"user_name" binding:"required,min=2,max=100"`
	Birth_date    time.Time `json:"birth_date"`
	Email         string    `json:"email" binding:"required,email"`
	Phone_num     string    `json:"phone_num" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	Password_hash string    `json:"password_hash" binding:"required"`
	Avatar_url    string    `json:"avatar_url" binding:"omitempty,url"`
}

type UpdateUsersPutRequest struct {
	Id_status  int32     `json:"id_status" binding:"required"`
	User_name  string    `json:"user_name" binding:"required,min=2,max=100"`
	Birth_date time.Time `json:"birth_date"`
	Email      string    `json:"email" binding:"required,email"`
	Phone_num  string    `json:"phone_num" binding:"required"`
	Address    string    `json:"address" binding:"required"`
}

type UpdateUsersPatchRequest struct {
	Id_status  *int32     `json:"id_status" binding:"omitempty,required"`
	User_name  *string    `json:"user_name" binding:"omitempty,required,min=2,max=100"`
	Birth_date *time.Time `json:"birth_date"`
	Email      *string    `json:"email" binding:"omitempty,required,email"`
	Phone_num  *string    `json:"phone_num" binding:"omitempty,required"`
	Address    *string    `json:"address" binding:"omitempty,required"`
}
