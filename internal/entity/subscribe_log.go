package entity

type SubscribeLog struct {
	ID int64 `gorm:"column:id primaryKey"` // id
	UserID int64 `gorm:"column:user_id"` // 用戶id
	TransactionID string `gorm:"column:transaction_id"` // 交易 id
	PurchaseDate string `gorm:"column:purchase_date"` // 訂閱購買日期
	ExpiresDate string `gorm:"column:expires_date"` // 訂閱過期日期
	Type string `gorm:"column:type"` // 紀錄類型(初次訂閱:initial_buy/恢復訂閱:resubscribe/續訂:renew/訂閱升級:upgrade/訂閱降級:downgrade/訂閱過期:expired/退費:refund)
	Msg string `gorm:"column:msg"` // 紀錄訊息
	CreateAt string `gorm:"column:create_at"` // 創建時間
}

func (SubscribeLog) TableName() string {
	return "subscribe_logs"
}