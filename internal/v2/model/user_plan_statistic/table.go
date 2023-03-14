package user_plan_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_plan_statistic/optional"

type Table struct {
	optional.UserIDField
	optional.PlanIDField
	optional.FinishWorkoutCountField
	optional.DurationField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "user_plan_statistics"
}
