package user_course_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_course_statistic/optional"

type Table struct {
	optional.UserIDField
	optional.CourseIDField
	optional.FinishWorkoutCountField
	optional.TotalFinishWorkoutCountField
	optional.DurationField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "user_course_statistics"
}
