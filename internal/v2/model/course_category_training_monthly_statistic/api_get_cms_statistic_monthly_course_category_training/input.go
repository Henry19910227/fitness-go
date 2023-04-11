package api_get_cms_statistic_monthly_course_category_training

import "github.com/Henry19910227/fitness-go/internal/v2/field/course_category_training_monthly_statistic/required"

// Input /v2/cms/statistic_monthly/course_category/training [GET]
type Input struct {
	Query Query
}
type Query struct {
	required.CategoryField
	required.YearField
	required.MonthField
}
