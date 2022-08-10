package receipt

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` // 收據id
}

type OrderIDOptional struct {
	OrderID *string `json:"order_id,omitempty" binding:"omitempty" example:"20220215104747115283"` //訂單id
}

type TransactionIDOptional struct {
	TransactionID *string `json:"transaction_id,omitempty" binding:"omitempty" example:"1000000968276600"` // 交易id
}

type OriginalTransactionIDOptional struct {
	OriginalTransactionID *string `json:"original_transaction_id,omitempty" binding:"omitempty" example:"1000000968276600"` // 初始交易id
}
