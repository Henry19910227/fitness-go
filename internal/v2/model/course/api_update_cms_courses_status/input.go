package api_update_cms_courses_status

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	updateLogOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course_status_update_log/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/cms/courses/course_status [PATCH]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	required.CourseStatusField
	updateLogOptional.CommentField
}
