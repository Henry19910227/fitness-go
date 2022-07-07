package order_subscribe_plan

import "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"

type Output struct {
	Table
	SubscribePlan *subscribe_plan.Output `json:"subscribe_plan,omitempty" gorm:"foreignKey:id;references:subscribe_plan_id"`
}

func (Output) TableName() string {
	return "order_subscribe_plans"
}
