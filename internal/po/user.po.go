package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Bảng User
type User struct {
	gorm.Model  // Sẽ bao gồm các trường ID, CreatedAt, UpdatedAt và DeletedAt
	UUID uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; unique"`	// field UUID là khóa chính
	Name string `gorm:"column:name; type:varchar(255); not null"`			// field Name là tên người dùng, không được để trống
	Email string `gorm:"column:email; type:varchar(255); not null; unique"`	// field Email là email người dùng, không được để trống và phải là duy nhất
	Password string `gorm:"column:password; type:varchar(255); not null"`		// field Password là mật khẩu người dùng, không được để trống
	IsActive bool `gorm:"column:is_active; type:boolean;"`
	Roles []Role `gorm:"many2many:go_user_roles;"`
}

// TableName sẽ trả về tên bảng tương ứng với struct User
func(u *User) TableName() string {
	return "go_db_users"
}