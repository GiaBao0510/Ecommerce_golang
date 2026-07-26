package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToUserModel(u database.USER) models.Users {
	return models.Users{
		Uuid:          u.Uuid,
		User_name:     u.UserName,
		Id_status:     u.IDStatus.Int32,
		Birth_date:    u.DateOfBirth.Time,
		Email:         u.Email,
		Phone_num:     u.PhoneNum.String,
		Address:       u.Address.String,
		Password_hash: u.PasswordHash,
		Avatar_url:    u.AvatarUrl.String,
		Created_at:    u.CreatedAt.Time,
		Updated_at:    u.UpdatedAt,
		Deleted_at:    u.DeletedAt,
		Is_email_verified: u.IsEmailVerified.Bool,
		Is_phone_verified: u.IsPhoneVerified.Bool,
	}
}
