package api_create_favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/favorite_course/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/favorite/course/{course_id} [POST]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.CourseIDField
}
