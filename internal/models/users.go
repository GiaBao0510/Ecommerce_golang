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
