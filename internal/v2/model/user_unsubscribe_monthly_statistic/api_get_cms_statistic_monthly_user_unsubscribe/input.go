package api_get_cms_statistic_monthly_user_unsubscribe

import "github.com/Henry19910227/fitness-go/internal/v2/field/user_unsubscribe_monthly_statistic/optional"

// Input /v2/cms/statistic_monthly/user/unsubscribe [GET]
type Input struct {
	Query Query
}
type Query struct {
	optional.YearField
	optional.MonthField
}
