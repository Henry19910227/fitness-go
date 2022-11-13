package optional

type IDField struct {
	ID *string `json:"id,omitempty" form:"order_id" gorm:"column:id" binding:"omitempty" example:"202105201300687423"` //訂單id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" form:"user_id" binding:"omitempty" example:"10001"` //用戶id
}
type QuantityField struct {
	Quantity *int `json:"quantity,omitempty" gorm:"column:quantity" binding:"omitempty" example:"1"` //數量
}
type OrderTypeField struct {
	Type *int `json:"order_type,omitempty" gorm:"column:order_type" form:"type" binding:"omitempty,oneof=1 2" example:"1"` //訂單類型(1:課表購買/2:會員訂閱)
}
type OrderStatusField struct {
	OrderStatus *int `json:"order_status,omitempty" gorm:"column:order_status" form:"order_status" binding:"omitempty,oneof=1 2 3 4" example:"2"` //訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
