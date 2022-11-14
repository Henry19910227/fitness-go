package api_get_trainer_course_overview

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/trainer/course/{course_id}/overview [GET]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.CourseIDField
}
