package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		d.Time = time.Time{}
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value == "" {
		d.Time = time.Time{}
		return nil
	}

	parsed, err := time.Parse("02-01-2006", value)
	if err != nil {
		return fmt.Errorf("birth_date phải có định dạng dd-mm-yyyy: %w", err)
	}

	d.Time = parsed
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(d.Time.Format("02-01-2006"))
}

type Users struct {
	Uuid                 string `json:"uuid"`
	Id_status            int32  `json:"id_status"`
	User_name            string `json:"user_name"`
	Birth_date           time.Time
	Email                string `json:"email"`
	Phone_num            string `json:"phone_num"`
	Is_email_verified    bool   `json:"is_email_verified"`
	Is_phonenum_verified bool   `json:"is_phonenum_verified"`
	Address              string `json:"address"`
	Password_hash        string `json:"password_hash"`
	Avatar_url           string `json:"avatar_url"`
	Created_at           time.Time
	Updated_at           sql.NullTime
	Deleted_at           sql.NullTime
}

type CreateUsersRequest struct {
	Uuid          string   // Thuộc tính này sẽ được thư viện uuid tự động tạo ra, nên không cần binding
	Id_status     int32    `json:"id_status" binding:"required"`
	User_name     string   `json:"user_name" binding:"required,min=2,max=100"`
	Birth_date    DateOnly `json:"birth_date"`
	Email         string   `json:"email" binding:"required,email"`
	Phone_num     string   `json:"phone_num" binding:"required"`
	Address       string   `json:"address" binding:"required"`
	Password_hash string   `json:"password_hash" binding:"required"`
	Avatar_url    string   `json:"avatar_url" binding:"omitempty,url"`
}

type RegisterRequest struct {
	Uuid          string   // Thuộc tính này sẽ được thư viện uuid tự động tạo ra, nên không cần binding
	Id_status     int32    `json:"id_status" binding:"required"`
	User_name     string   `json:"user_name" binding:"required,min=2,max=100"`
	Birth_date    DateOnly `json:"birth_date"`
	Email         string   `json:"email" binding:"required,email"`
	Phone_num     string   `json:"phone_num" binding:"required"`
	Address       string   `json:"address" binding:"required"`
	Password_hash string   `json:"password_hash" binding:"required"`
	Avatar_url    string   `json:"avatar_url" binding:"omitempty,url"`
}

type UpdateUsersPutRequest struct {
	Id_status  int32    `json:"id_status" binding:"required"`
	User_name  string   `json:"user_name" binding:"required,min=2,max=100"`
	Birth_date DateOnly `json:"birth_date"`
	Email      string   `json:"email" binding:"required,email"`
	Phone_num  string   `json:"phone_num" binding:"required"`
	Address    string   `json:"address" binding:"required"`
}

type UpdateUsersPatchRequest struct {
	Id_status  *int32    `json:"id_status" binding:"omitempty,required"`
	User_name  *string   `json:"user_name" binding:"omitempty,required,min=2,max=100"`
	Birth_date *DateOnly `json:"birth_date"`
	Email      *string   `json:"email" binding:"omitempty,required,email"`
	Phone_num  *string   `json:"phone_num" binding:"omitempty,required"`
	Address    *string   `json:"address" binding:"omitempty,required"`
}

type LoginRequest struct {
	Account   string `json:"account" binding:"required"` // Có thể là email hoặc số điện thoại
	Passoword string `json:"password" binding:"required"`
}

type LoginInfo struct {
	Email         string `json:"email"`
	Role_id       int32  `json:"role_id"`
	Password_hash string `json:"password_hash"`
	Id_status     int32  `json:"id_status"`
}
