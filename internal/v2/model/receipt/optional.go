package receipt

type OrderIDOptional struct {
	OrderID *string `json:"order_id,omitempty" binding:"omitempty" example:"20220215104747115283"` //訂單id
}
