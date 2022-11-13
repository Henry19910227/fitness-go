package optional

type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" uri:"order_id" form:"order_id" binding:"omitempty" example:"202105201300687423"` //訂單id
}
