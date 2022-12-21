package api_get_trainer_course_statistics

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	courseUsageStatisticRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course_usage_statistic/required"

	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

// Output /v2/trainer/course/statistics [GET]
type Output struct {
	base.Output
	Data   Data           `json:"data,omitempty"`
	Paging *paging.Output `json:"paging,omitempty"`
}
type Data []*struct {
	courseOptional.IDField
	courseOptional.SaleTypeField
	courseOptional.SaleIDField
	courseOptional.CourseStatusField
	courseOptional.CategoryField
	courseOptional.ScheduleTypeField
	courseOptional.NameField
	courseOptional.CoverField
	courseOptional.LevelField
	courseOptional.PlanCountField
	courseOptional.WorkoutCountField
	courseOptional.CreateAtField
	courseOptional.UpdateAtField
	CourseUsageStatistic struct {
		courseUsageStatisticRequired.TotalFinishWorkoutCountField
		courseUsageStatisticRequired.UserFinishCountField
		courseUsageStatisticRequired.FinishCountAvgField
	} `json:"course_usage_statistic"`
}
