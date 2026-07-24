package dto

type UserResponseBase struct {
	Uuid          string `json:"uuid"`
	Password_hash string `json:"password_hash"`
}
