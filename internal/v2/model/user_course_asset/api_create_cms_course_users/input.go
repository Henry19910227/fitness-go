package api_create_cms_course_users

import courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"

// Input /v2/cms/course/{course_id}/users [POST]
type Input struct {
	Uri  Uri
	Body Body
}
type Uri struct {
	courseRequired.CourseIDField
}
type Body struct {
	UserIDs []int64 `json:"user_ids" binding:"required"` // 用戶id列表
}
