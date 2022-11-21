package optional

type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" binding:"omitempty" example:"20220215104747115283"` // 訂單id
}
type SubscribePlanIDField struct {
	SubscribePlanID *int64 `json:"subscribe_plan_id,omitempty" gorm:"column:subscribe_plan_id" binding:"omitempty" example:"1"` // 訂閱方案id
}
type StatusField struct {
	Status *int `json:"status,omitempty" gorm:"column:status" binding:"omitempty" example:"1"` // 訂閱狀態(0:無訂閱/1:訂閱中)
}
