package order_subscribe_plan

import "github.com/Henry19910227/fitness-go/internal/v2/field/order_subscribe_plan/optional"

type Table struct {
	optional.OrderIDField
	optional.SubscribePlanIDField
	optional.StatusField
}

func (Table) TableName() string {
	return "order_subscribe_plans"
}
