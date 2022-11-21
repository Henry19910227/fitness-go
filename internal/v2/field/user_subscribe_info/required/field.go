package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type OrderIDField struct {
	OrderID string `json:"order_id" gorm:"column:order_id" binding:"required" example:"20220215104747115283"` //訂單id
}
type OriginalTransactionIDField struct {
	OriginalTransactionID string `json:"original_transaction_id" gorm:"column:original_transaction_id" binding:"required" example:"1000000968276600"` // 初始交易id
}
type StatusField struct {
	Status int `json:"status" gorm:"column:status" binding:"required" example:"1"` // 會員狀態(0:無會員狀態/1:付費會員狀態)
}
type PaymentTypeField struct {
	PaymentType int `json:"payment_type" gorm:"column:payment_type" binding:"required" example:"1"` //支付方式(0:尚未指定/1:apple內購/2:google內購)
}
type StartDateField struct {
	StartDate string `json:"start_date" gorm:"column:start_date" binding:"required" example:"2022-07-11 11:00:00"` // 訂閱開始日期
}
type ExpiresDateField struct {
	ExpiresDate string `json:"end_date" gorm:"column:expires_date" binding:"required" example:"2023-07-11 11:00:00"` // 訂閱結束日期
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
