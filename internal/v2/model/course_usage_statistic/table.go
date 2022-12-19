package course_usage_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_usage_statistic/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.TotalFinishWorkoutCountField
	optional.UserFinishCountField
	optional.FinishCountAvgField
	optional.MaleFinishCountField
	optional.FemaleFinishCountField
	optional.Age13to17CountField
	optional.Age18to24CountField
	optional.Age25to34CountField
	optional.Age35to44CountField
	optional.Age45to54CountField
	optional.Age55to64CountField
	optional.Age65upCountField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "course_usage_statistics"
}
