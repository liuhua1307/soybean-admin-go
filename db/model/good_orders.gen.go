// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGoodOrder = "good_orders"

// GoodOrder mapped from table <good_orders>
type GoodOrder struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID  int64  `gorm:"column:order_id;not null" json:"order_id"`
	GoodID   int64  `gorm:"column:good_id;not null" json:"good_id"`
	Quantity int32  `gorm:"column:quantity;not null" json:"quantity"`
	Order    *Order `gorm:"foreignKey:order_id" json:"order"`
	Good     *Good  `gorm:"foreignKey:good_id" json:"good"`
}

// TableName GoodOrder's table name
func (*GoodOrder) TableName() string {
	return TableNameGoodOrder
}
