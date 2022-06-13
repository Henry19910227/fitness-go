package entity

type OrderSubscribePlan struct {
	OrderID         string `gorm:"column:order_id"`          // 訂單id
	SubscribePlanID int64  `gorm:"column:subscribe_plan_id"` // 訂閱方案id
}

func (OrderSubscribePlan) TableName() string {
	return "order_subscribe_plans"
}

type OrderSubscribePlanTemplate struct {
	OrderID         string `gorm:"column:order_id"`          // 訂單id
	SubscribePlanID int64  `gorm:"column:subscribe_plan_id"` // 訂閱方案id
}

func (OrderSubscribePlanTemplate) TableName() string {
	return "order_subscribe_plans"
}
