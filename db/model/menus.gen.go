// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMenu = "menus"

// Menu mapped from table <menus>
type Menu struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateBy   string `gorm:"column:create_by;comment:创建者名称" json:"create_by"`                              // 创建者名称
	UpdateBy   string `gorm:"column:update_by;comment:修改者名称" json:"update_by"`                              // 修改者名称
	Status     string `gorm:"column:status;default:1;comment:状态（1-启用，2-禁用）" json:"status"`                  // 状态（1-启用，2-禁用）
	ParentID   int64  `gorm:"column:parent_id;default:1;comment:状态（1-启用，2-禁用）" json:"parent_id"`            // 状态（1-启用，2-禁用）
	MenuType   string `gorm:"column:menu_type;default:1;comment:类型（1-目录，2-菜单）" json:"menu_type"`            // 类型（1-目录，2-菜单）
	MenuName   string `gorm:"column:menu_name;not null;comment:菜单名称" json:"menu_name"`                      // 菜单名称
	RouteName  string `gorm:"column:route_name;not null;comment:路由名称" json:"route_name"`                    // 路由名称
	RoutePath  string `gorm:"column:route_path;not null;comment:路由路径" json:"route_path"`                    // 路由路径
	Component  string `gorm:"column:component;comment:组件路径" json:"component"`                               // 组件路径
	Order      int64  `gorm:"column:order;comment:排序" json:"order"`                                         // 排序
	I18nKey    string `gorm:"column:i18n_key;comment:排序" json:"i18n_key"`                                   // 排序
	Icon       string `gorm:"column:icon;comment:Icon" json:"icon"`                                         // Icon
	IconType   string `gorm:"column:icon_type;default:1;comment:图标类型（1-iconify图标，2-本地图标）" json:"icon_type"` // 图标类型（1-iconify图标，2-本地图标）
	HideInMenu bool   `gorm:"column:hide_in_menu;comment:隐藏菜单" json:"hide_in_menu"`                         // 隐藏菜单
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UpdateTime string `gorm:"column:update_time" json:"update_time"`
	Roles      []Role `gorm:"many2many:roles_menus" json:"roles"`
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
