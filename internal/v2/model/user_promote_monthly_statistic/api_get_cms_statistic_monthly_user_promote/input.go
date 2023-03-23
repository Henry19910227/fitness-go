package api_get_cms_statistic_monthly_user_promote

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_promote_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/user/promote [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.YearField
	required.MonthField
}
