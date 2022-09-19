package course_training_avg_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_training_avg_statistic/optional"

type Table struct {
	optional.IDField
	optional.CourseIDField
	optional.RateField
	optional.CreateAtField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "course_training_avg_statistics"
}
