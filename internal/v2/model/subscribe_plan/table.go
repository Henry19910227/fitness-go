package subscribe_plan

import "github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_plan/optional"

type Table struct {
	optional.IDField
	optional.ProductLabelIDField
	optional.PeriodField
	optional.NameField
	optional.EnableField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "subscribe_plans"
}
