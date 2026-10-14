package models

import (
	"database/sql"
)

type ProductDetail struct {
	IDProduct      string       `json:"id_product" validate:"required"`
	Price          string       `json:"price" validate:"required"`
	LikeCount      int32        `json:"like_count"`
	Quantity       int32        `json:"quantity"`
	ProductionDate sql.NullTime `json:"production_date"`
	ExpirationDate sql.NullTime `json:"expiration_date"`
	UpdatedAt      sql.NullTime `json:"updated_at"`
}
