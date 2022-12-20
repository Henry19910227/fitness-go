package user_course_usage_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_course_usage_monthly_statistic/optional"

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.FreeUsageCountField
	optional.SubscribeUsageCountField
	optional.ChargeUsageCountField
	optional.YearField
	optional.MonthField
	optional.CreateAtField
	optional.UpdateAtField
}
func (Table) TableName() string {
	return "user_course_usage_monthly_statistics"
}
