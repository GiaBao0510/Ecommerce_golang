package po

import (
	"gorm.io/gorm"
)

// Bảng Role
type Role struct {
	gorm.Model
	ID int64 `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'Primary Key is ID'"`	
	RoleName string `gorm:"column:role_name; type:varchar(255); not null"`			
	RoleNote string `gorm:"column:role_note; type:text;"`	
}

// TableName sẽ trả về tên bảng tương ứng với struct Role
func(r *Role) TableName() string {
	return "go_db_roles"
}