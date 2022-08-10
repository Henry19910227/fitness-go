package order_subscribe_plan

type OrderIDOptional struct {
	OrderID *string `json:"order_id,omitempty" binding:"omitempty" example:"20220215104747115283"` // 訂單id
}
type SubscribePlanIDOptional struct {
	SubscribePlanID *int64 `json:"subscribe_plan_id,omitempty" binding:"omitempty" example:"1"` // 訂閱方案id
}
