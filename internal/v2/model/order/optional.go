package order

type IDOptional struct {
	ID *string `json:"id,omitempty" form:"order_id" binding:"omitempty" example:"202105201300687423"` //訂單id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" form:"user_id" binding:"omitempty" example:"10001"` //用戶id
}
type QuantityOptional struct {
	Quantity *int `json:"quantity,omitempty" binding:"omitempty" example:"1"` //數量
}
type OrderTypeOptional struct {
	Type *int `json:"order_type,omitempty" form:"type" binding:"omitempty,oneof=1 2" example:"1"` //訂單類型(1:課表購買/2:會員訂閱)
}
type OrderStatusOptional struct {
	OrderStatus *int `json:"order_status,omitempty" form:"order_status" binding:"omitempty,oneof=1 2 3 4" example:"2"` //訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
type CreateAtOptional struct {
	CreateAt *string `json:"create_at,omitempty" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtOptional struct {
	UpdateAt *string `json:"update_at,omitempty" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
