// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRoleMenu = "role_menus"

// RoleMenu mapped from table <role_menus>
type RoleMenu struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID     int64     `gorm:"column:role_id;not null;comment:角色id" json:"role_id"` // 角色id
	MenuID     int64     `gorm:"column:menu_id;not null;comment:菜单id" json:"menu_id"` // 菜单id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Role       *Role     `json:"role"`
	Menu       *Menu     `json:"menu"`
}

// TableName RoleMenu's table name
func (*RoleMenu) TableName() string {
	return TableNameRoleMenu
}
