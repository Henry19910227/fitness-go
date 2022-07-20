package order

type IDRequired struct {
	ID string `json:"id" form:"order_id" binding:"required" example:"202105201300687423"` //訂單id
}
type UserIDRequired struct {
	UserID int64 `json:"user_id" form:"user_id" binding:"required" example:"10001"` //用戶id
}
type QuantityRequired struct {
	Quantity int `json:"quantity" binding:"required" example:"1"` //數量
}
type OrderTypeRequired struct {
	Type int `json:"order_type" form:"type" binding:"required,oneof=1 2" example:"1"` //訂單類型(1:課表購買/2:會員訂閱)
}
type OrderStatusRequired struct {
	OrderStatus int `json:"order_status" form:"order_status" binding:"required,oneof=1 2 3 4" example:"2"` //訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
type CreateAtRequired struct {
	CreateAt string `json:"create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtRequired struct {
	UpdateAt string `json:"update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
