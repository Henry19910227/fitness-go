package required

type IDField struct {
	ID string `json:"id" form:"order_id" gorm:"column:id" binding:"required" example:"202105201300687423"` //訂單id
}
type UserIDField struct {
	UserID int64 `json:"user_id" form:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type QuantityField struct {
	Quantity int `json:"quantity" gorm:"column:quantity" binding:"required" example:"1"` //數量
}
type OrderTypeField struct {
	Type int `json:"order_type" form:"type" gorm:"column:order_type" binding:"required,oneof=1 2" example:"1"` //訂單類型(1:課表購買/2:會員訂閱)
}
type OrderStatusField struct {
	OrderStatus int `json:"order_status" gorm:"column:order_status" form:"order_status" binding:"required,oneof=1 2 3 4" example:"2"` //訂單狀態(1:等待付款/2:已付款/3:錯誤/4:取消)
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
