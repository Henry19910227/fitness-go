package model

type Receipt struct {
	ID                    int64  `gorm:"column:id"`                      // 收據id
	OrderID               string `gorm:"column:order_id"`                // 訂單id
	PaymentType           int    `gorm:"column:payment_type"`            // 支付方式(1:apple內購/2:google內購)
	ReceiptToken          string `gorm:"column:receipt_token"`           // 收據token
	OriginalTransactionID string `gorm:"column:original_transaction_id"` // 初始交易id
	TransactionID         string `gorm:"column:transaction_id"`          // 交易id
	ProductID             string `gorm:"column:product_id"`              // 產品id
	Quantity              int    `gorm:"column:quantity"`                // 購買數量
	CreateAt              string `gorm:"column:create_at"`               // 創建日期
}

type CreateReceiptParam struct {
	OrderID               string `gorm:"column:order_id"`                // 訂單id
	PaymentType           int    `gorm:"column:payment_type"`            // 支付方式(1:apple內購/2:google內購)
	ReceiptToken          string `gorm:"column:receipt_token"`           // 收據token
	OriginalTransactionID string `gorm:"column:original_transaction_id"` // 初始交易id
	TransactionID         string `gorm:"column:transaction_id"`          // 交易id
	ProductID             string `gorm:"column:product_id"`              // 產品id
	Quantity              int    `gorm:"column:quantity"`                // 購買數量
}
