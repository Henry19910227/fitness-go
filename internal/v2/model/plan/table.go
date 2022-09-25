package plan

import "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.NameField
	optional.WorkoutCountField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "plans"
}
