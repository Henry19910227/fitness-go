package api_get_cms_user_subscribe_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/user/subscribe [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.YearField
	required.MonthField
}
