package api_get_trainer_course_statistic

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	courseUsageStatisticOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course_usage_statistic/optional"
	reviewRequired "github.com/Henry19910227/fitness-go/internal/v2/field/review_statistic/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/course/{course_id}/statistic [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
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
	ReviewStatistic struct {
		reviewRequired.ScoreTotalField
		reviewRequired.AmountField
	} `json:"review_statistic"`
	CourseUsageStatistic struct {
		courseUsageStatisticOptional.TotalFinishWorkoutCountField
		courseUsageStatisticOptional.UserFinishCountField
		courseUsageStatisticOptional.FinishCountAvgField
		courseUsageStatisticOptional.MaleFinishCountField
		courseUsageStatisticOptional.FemaleFinishCountField
		courseUsageStatisticOptional.Age13to17CountField
		courseUsageStatisticOptional.Age18to24CountField
		courseUsageStatisticOptional.Age25to34CountField
		courseUsageStatisticOptional.Age35to44CountField
		courseUsageStatisticOptional.Age45to54CountField
		courseUsageStatisticOptional.Age55to64CountField
		courseUsageStatisticOptional.Age65upCountField
		courseUsageStatisticOptional.CreateAtField
		courseUsageStatisticOptional.UpdateAtField
	} `json:"course_usage_statistic"`
}
