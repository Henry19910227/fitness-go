package api_get_store_course

import courseRequired "github.com/Henry19910227/fitness-go/internal/v2/field/course/required"

// Input /v2/store/course/{course_id} [GET]
type Input struct {
	courseRequired.UserIDField
	Uri APIGetStoreCourseUri
}
type APIGetStoreCourseUri struct {
	courseRequired.IDField
}
