package subscribe_log

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` // log id
}
type OriginalTransactionIDField struct {
	OriginalTransactionID *string `json:"original_transaction_id,omitempty" gorm:"column:original_transaction_id" example:"1000000968276600"` // 初始交易id
}
type TransactionIDField struct {
	TransactionID *string `json:"transaction_id,omitempty" gorm:"column:transaction_id" example:"1000000968276600"` // 交易id
}
type PurchaseDateField struct {
	PurchaseDate *string `json:"purchase_date,omitempty" gorm:"column:purchase_date" example:"2022-06-14 00:00:00"` // 訂閱購買日期
}
type ExpiresDateField struct {
	ExpiresDate *string `json:"expires_date,omitempty" gorm:"column:expires_date" example:"2022-06-14 00:00:00"` //訂閱過期日期
}
type TypeField struct {
	Type *string `json:"type,omitempty" gorm:"column:type" example:"3"` // 紀錄類型(初次訂閱:initial_buy/恢復訂閱:resubscribe/續訂:renew/訂閱升級:upgrade/訂閱降級:downgrade/訂閱過期:expired/退費:refund)
}
type MsgField struct {
	Msg *string `json:"msg,omitempty" gorm:"column:msg" example:"hello world"` // 紀錄訊息
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	OriginalTransactionIDField
	TransactionIDField
	PurchaseDateField
	ExpiresDateField
	TypeField
	MsgField
	CreateAtField
}

func (Table) TableName() string {
	return "subscribe_logs"
}
