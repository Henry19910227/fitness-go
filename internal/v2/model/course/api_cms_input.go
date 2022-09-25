package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
	"mime/multipart"
)

// APIGetCMSCoursesInput /v2/cms/courses [GET]
type APIGetCMSCoursesInput struct {
	optional.IDField
	optional.NameField
	optional.CourseStatusField
	optional.SaleTypeField
	PagingInput
	OrderByInput
}

// APIGetCMSCourseInput /v2/cms/course/{course_id} [GET]
type APIGetCMSCourseInput struct {
	required.IDField
}

// APIUpdateCMSCoursesStatusInput /v2/cms/courses/course_status [PATCH]
type APIUpdateCMSCoursesStatusInput struct {
	IDs []int64 `json:"ids" binding:"required"` // 課表 id
	required.CourseStatusField
}

// APIUpdateCMSCourseCoverInput /v2/cms/course/{course_id}/cover [PATCH]
type APIUpdateCMSCourseCoverInput struct {
	required.IDField
	CoverNamed string
	File       multipart.File
}
