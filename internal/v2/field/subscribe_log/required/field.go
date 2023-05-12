package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` // log id
}
type OriginalTransactionIDField struct {
	OriginalTransactionID string `json:"original_transaction_id" gorm:"column:original_transaction_id" binding:"required" example:"1000000968276600"` // 初始交易id
}
type TransactionIDField struct {
	TransactionID string `json:"transaction_id" gorm:"column:transaction_id" binding:"required" example:"1000000968276600"` // 交易id
}
type PurchaseDateField struct {
	PurchaseDate string `json:"purchase_date" gorm:"column:purchase_date" binding:"required" example:"2022-06-14 00:00:00"` // 訂閱購買日期
}
type ExpiresDateField struct {
	ExpiresDate string `json:"expires_date" gorm:"column:expires_date" binding:"required" example:"2022-06-14 00:00:00"` //訂閱過期日期
}
type TypeField struct {
	Type string `json:"type" gorm:"column:type" binding:"required" example:"3"` // 紀錄類型(初次訂閱:initial_buy/恢復訂閱:resubscribe/續訂:renew/訂閱升級:upgrade/訂閱降級:downgrade/訂閱過期:expired/退費:refund)
}
type MsgField struct {
	Msg string `json:"msg" gorm:"column:msg" binding:"required" example:"hello world"` // 紀錄訊息
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
