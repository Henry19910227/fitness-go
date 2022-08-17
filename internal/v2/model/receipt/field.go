package receipt

type IDField struct {
	ID *int64 ` json:"id,omitempty" gorm:"column:id" example:"1"` // 收據id
}
type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" example:"20220215104747115283"` //訂單id
}
type PaymentTypeField struct {
	PaymentType *int `json:"payment_type,omitempty" gorm:"column:payment_type" example:"1"` //支付方式(0:尚未指定/1:apple內購/2:google內購)
}
type ReceiptTokenField struct {
	ReceiptToken *string `json:"receipt_token,omitempty" gorm:"column:receipt_token" example:"XXXX"` // 收據token
}
type OriginalTransactionIDField struct {
	OriginalTransactionID *string `json:"original_transaction_id,omitempty" gorm:"column:original_transaction_id" example:"1000000968276600"` // 初始交易id
}
type TransactionIDField struct {
	TransactionID *string `json:"transaction_id,omitempty" gorm:"column:transaction_id" example:"1000000968276600"` // 交易id
}
type ProductIDField struct {
	ProductID *string `json:"product_id,omitempty" gorm:"column:product_id" example:"com.fitness.gold_member_month"` // 產品id
}
type QuantityField struct {
	Quantity *int `json:"quantity,omitempty" gorm:"column:quantity" example:"1"` //數量
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	OrderIDField
	PaymentTypeField
	ReceiptTokenField
	OriginalTransactionIDField
	TransactionIDField
	ProductIDField
	QuantityField
	CreateAtField
}

func (Table) TableName() string {
	return "receipts"
}
