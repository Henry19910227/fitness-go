package order_subscribe_plan

type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" example:"20220215104747115283"` // 訂單id
}
type SubscribePlanIDField struct {
	SubscribePlanID *int64 `json:"subscribe_plan_id,omitempty" gorm:"column:subscribe_plan_id" example:"1"` // 訂閱方案id
}

type Table struct {
	OrderIDField
	SubscribePlanIDField
}

func (Table) TableName() string {
	return "order_subscribe_plans"
}
