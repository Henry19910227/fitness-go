package api_get_trainer_course_usage_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_usage_monthly_statistic/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_course_usage_monthly_statistic/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/course/usage_monthly_statistic [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	required.FreeUsageCountField
	required.SubscribeUsageCountField
	required.ChargeUsageCountField
	optional.YearField
	optional.MonthField
	optional.CreateAtField
	optional.UpdateAtField
}
