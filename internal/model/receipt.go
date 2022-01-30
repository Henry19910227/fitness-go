package model

type CreateReceiptParam struct {
	OrderID               string `gorm:"column:order_id"`                // 訂單id
	PaymentType           int    `gorm:"column:payment_type"`            // 支付方式(1:apple內購/2:google內購)
	ReceiptToken          string `gorm:"column:receipt_token"`           // 收據token
	OriginalTransactionID string `gorm:"column:original_transaction_id"` // 初始交易id
	TransactionID         string `gorm:"column:transaction_id"`          // 交易id
	ProductID             string `gorm:"column:product_id"`              // 產品id
	Quantity              int    `gorm:"column:quantity"`                // 購買數量
}
