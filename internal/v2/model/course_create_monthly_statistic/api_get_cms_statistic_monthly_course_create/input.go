package api_get_cms_statistic_monthly_course_create

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_create_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/course/create [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.YearField
	required.MonthField
}
