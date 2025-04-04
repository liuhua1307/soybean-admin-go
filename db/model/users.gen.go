// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateBy   string    `gorm:"column:create_by;comment:创建者名称" json:"create_by"`                     // 创建者名称
	UpdateBy   string    `gorm:"column:update_by;comment:修改者名称" json:"update_by"`                     // 修改者名称
	Status     int64     `gorm:"column:status;default:1;comment:状态（1-正常，2-禁用）" json:"status"`         // 状态（1-正常，2-禁用）
	UserName   string    `gorm:"column:user_name;not null;comment:用户名" json:"user_name"`              // 用户名
	UserGender string    `gorm:"column:user_gender;default:1;comment:性别（1-男，2-女）" json:"user_gender"` // 性别（1-男，2-女）
	NickName   string    `gorm:"column:nick_name;comment:昵称" json:"nick_name"`                        // 昵称
	UserPhone  string    `gorm:"column:user_phone;comment:电话" json:"user_phone"`                      // 电话
	UserEmail  string    `gorm:"column:user_email;comment:邮箱" json:"user_email"`                      // 邮箱
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	Password   string    `gorm:"column:password" json:"password"`
	Roles      []Role    `gorm:"many2many:user_roles" json:"roles"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
