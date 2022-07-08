package receipt

type OrderIDRequired struct {
	OrderID string `json:"order_id" uri:"order_id" binding:"required" example:"20220215104747115283"` //訂單id
}
