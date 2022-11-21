package required

type OrderIDField struct {
	OrderID string `json:"order_id" gorm:"column:order_id" binding:"required" example:"20220215104747115283"` // 訂單id
}
type SubscribePlanIDField struct {
	SubscribePlanID int64 `json:"subscribe_plan_id" gorm:"column:subscribe_plan_id" binding:"required" example:"1"` // 訂閱方案id
}
type StatusField struct {
	Status int `json:"status" gorm:"column:status" binding:"required" example:"1"` // 訂閱狀態(0:無訂閱/1:訂閱中)
}
