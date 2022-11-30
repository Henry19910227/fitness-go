package api_get_trainer_course

import "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"

// Input /v2/trainer/course/{course_id} [GET]
type Input struct {
	required.UserIDField
	Uri Uri
}
type Uri struct {
	required.CourseIDField
}
