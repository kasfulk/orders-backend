// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrderItem = "order_items"

// OrderItem mapped from table <order_items>
type OrderItem struct {
	ID           int64   `gorm:"column:id;primaryKey" json:"id"`
	OrderID      *int64  `gorm:"column:order_id" json:"order_id"`
	PricePerUnit *string `gorm:"column:price_per_unit" json:"price_per_unit"`
	Quantity     *string `gorm:"column:quantity" json:"quantity"`
	Product      *string `gorm:"column:product" json:"product"`
}

// TableName OrderItem's table name
func (*OrderItem) TableName() string {
	return TableNameOrderItem
}
