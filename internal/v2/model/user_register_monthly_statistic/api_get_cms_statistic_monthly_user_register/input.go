package api_get_cms_statistic_monthly_user_register

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_register_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/user/register [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.YearField
	required.MonthField
}
