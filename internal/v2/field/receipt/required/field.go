package required

type IDField struct {
	ID int64 ` json:"id" gorm:"column:id" binding:"required" example:"1"` // 收據id
}
type OrderIDField struct {
	OrderID string `json:"order_id" uri:"order_id" gorm:"column:order_id" binding:"required" example:"20220215104747115283"` //訂單id
}
type PaymentTypeField struct {
	PaymentType int `json:"payment_type" gorm:"column:payment_type" binding:"required" example:"1"` //支付方式(0:尚未指定/1:apple內購/2:google內購)
}
type ReceiptTokenField struct {
	ReceiptToken string `json:"receipt_token" gorm:"column:receipt_token" binding:"required" example:"XXXX"` // 收據token
}
type OriginalTransactionIDField struct {
	OriginalTransactionID string `json:"original_transaction_id" gorm:"column:original_transaction_id" binding:"required" example:"1000000968276600"` // 初始交易id
}
type TransactionIDField struct {
	TransactionID string `json:"transaction_id" gorm:"column:transaction_id" binding:"required" example:"1000000968276600"` // 交易id
}
type ProductIDField struct {
	ProductID string `json:"product_id" gorm:"column:product_id" binding:"required" example:"com.fitness.gold_member_month"` // 產品id
}
type QuantityField struct {
	Quantity int `json:"quantity" gorm:"column:quantity" binding:"required" example:"1"` //數量
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
