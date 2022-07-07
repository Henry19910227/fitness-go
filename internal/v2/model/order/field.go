package order

type IDField struct {
	ID *string `json:"id,omitempty" gorm:"column:id" example:"202105201300687423"` //訂單id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type QuantityField struct {
	Quantity *int `json:"quantity,omitempty" gorm:"column:quantity" example:"1"` //數量
}
type OrderTypeField struct {
	Type *int `json:"order_type,omitempty" gorm:"column:order_type" example:"1"` //訂單類型(1:課表購買/2:會員訂閱)
}
type OrderStatusField struct {
	OrderStatus *int `json:"order_status,omitempty" gorm:"column:order_status" example:"2"` //訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	QuantityField
	OrderTypeField
	OrderStatusField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "foods"
}
