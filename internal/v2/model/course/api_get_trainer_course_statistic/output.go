package api_get_trainer_course_statistic

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	courseUsageStatisticRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course_usage_statistic/required"
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
		courseUsageStatisticRequired.TotalFinishWorkoutCountField
		courseUsageStatisticRequired.UserFinishCountField
		courseUsageStatisticRequired.FinishCountAvgField
		courseUsageStatisticRequired.MaleFinishCountField
		courseUsageStatisticRequired.FemaleFinishCountField
		courseUsageStatisticRequired.Age13to17CountField
		courseUsageStatisticRequired.Age18to24CountField
		courseUsageStatisticRequired.Age25to34CountField
		courseUsageStatisticRequired.Age35to44CountField
		courseUsageStatisticRequired.Age45to54CountField
		courseUsageStatisticRequired.Age55to64CountField
		courseUsageStatisticRequired.Age65upCountField
	} `json:"course_usage_statistic"`
}
