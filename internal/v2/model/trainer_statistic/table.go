package trainer_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/trainer_statistic/optional"

type Table struct {
	optional.UserIDField
	optional.StudentCountField
	optional.CourseCountField
	optional.ReviewScoreField
	optional.UpdateAtField
}

func (Table) TableName() string {
	return "trainer_statistics"
}
