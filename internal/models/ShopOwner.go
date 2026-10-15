package models

import (
	"database/sql"
	"time"
)

type ShopOwner struct {
	IDShopOwner int32        `json:"id_shop_owner"`
	Uuid        string       `json:"uuid" validate:"required"`
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"description"`
	AvatarUrl   string       `json:"avatar_url"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}
