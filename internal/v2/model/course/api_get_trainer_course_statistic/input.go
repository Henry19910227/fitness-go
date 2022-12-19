package api_get_trainer_course_statistic

import (
	courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/trainer/course/{course_id}/statistic [GET]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	courseRequired.CourseIDField
}
