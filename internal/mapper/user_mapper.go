package mapper

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
)

func ToUserModel(u database.User) models.Users {
	return models.Users{
		Uuid:                 u.Uuid,
		User_name:            u.UserName,
		Id_status:            u.IDStatus.Int32,
		Birth_date:           u.BirthDate.Time,
		Email:                u.Email,
		Phone_num:            u.PhoneNum.String,
		Address:              u.Address.String,
		Password_hash:        u.PasswordHash,
		Avatar_url:           u.AvatarUrl.String,
		Created_at:           u.CreatedAt.Time,
		Updated_at:           u.UpdatedAt,
		Deleted_at:           u.DeletedAt,
		Is_email_verified:    u.IsEmailVerified.Bool,
		Is_phonenum_verified: u.IsPhonenumVerified.Bool,
	}
}

func ToUserByRoleModel(ur database.GetUserByRoleIDRow) models.UserByRole {
	return models.UserByRole{
		Uuid:      ur.Uuid,
		User_name: ur.UserName,
		Email:     ur.Email,
		Phone_num: ur.PhoneNum.String,
		Address:   ur.Address.String,
	}
}

func ToRoleByUserModel(ru database.Role) models.RoleByUser {
	return models.RoleByUser{
		Role_id:     ru.RoleID,
		Role_name:   ru.RoleName,
		Description: ru.Description.String,
	}
}
