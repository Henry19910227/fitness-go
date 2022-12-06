package api_delete_cms_course_user

import (
	assetRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user_course_asset/required"
)

// Input /v2/cms/course/{course_id}/user/{user_id} [DELETE]
type Input struct {
	Uri Uri
}
type Uri struct {
	assetRequired.CourseIDField
	assetRequired.UserIDField
}
