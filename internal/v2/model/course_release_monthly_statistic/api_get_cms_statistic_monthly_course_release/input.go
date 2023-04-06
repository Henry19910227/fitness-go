package api_get_cms_statistic_monthly_course_release

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_release_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/course/release [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.YearField
	required.MonthField
}
