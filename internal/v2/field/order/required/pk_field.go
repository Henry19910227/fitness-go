package required

type OrderIDField struct {
	OrderID string `json:"order_id" uri:"order_id" form:"order_id" binding:"required" example:"202105201300687423"` //訂單id
}
