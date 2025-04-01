// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserRole = "user_roles"

// UserRole mapped from table <user_roles>
type UserRole struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID    int64     `gorm:"column:role_id;not null;comment:角色id" json:"role_id"` // 角色id
	UserID    int64     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"` // 用户id
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	User      *User     `gorm:"foreignKey:user_id" json:"user"`
	Role      *Role     `gorm:"foreignKey:role_id" json:"role"`
}

// TableName UserRole's table name
func (*UserRole) TableName() string {
	return TableNameUserRole
}
