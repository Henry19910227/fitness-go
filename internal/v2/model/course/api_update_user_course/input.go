package api_update_user_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
)

// Input /v2/user/course/{course_id} [PATCH]
type Input struct {
	required.UserIDField
	Uri  Uri
	Body Body
}
type Uri struct {
	required.CourseIDField
}
type Body struct {
	optional.NameField
}
