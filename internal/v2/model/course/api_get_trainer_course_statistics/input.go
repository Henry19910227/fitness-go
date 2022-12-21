package api_get_trainer_course_statistics

import (
	pageOptional "github.com/Henry19910227/fitness-go/internal/v2/field/paging/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/trainer/course/statistics [GET]
type Input struct {
	userRequired.UserIDField
	Query Query
}
type Query struct {
	pageOptional.PageField
	pageOptional.SizeField
}
