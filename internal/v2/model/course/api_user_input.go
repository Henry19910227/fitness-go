package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
)

// APIGetUserCoursesInput /v2/user/courses [GET]
type APIGetUserCoursesInput struct {
	required.UserIDField
	Query APIGetUserCoursesQuery
}
type APIGetUserCoursesQuery struct {
	Type int `form:"type" binding:"required,oneof=1 2 3" example:"1"` // 搜尋類別(1:進行中課表/2:付費課表/3:個人課表)
	PagingInput
}

// APICreateUserCourseInput /v2/user/course [POST]
type APICreateUserCourseInput struct {
	required.UserIDField
	Body APICreateUserCourseBody
}
type APICreateUserCourseBody struct {
	required.NameField
	required.ScheduleTypeField
}

// APIDeleteUserCourseInput /v2/user/course/{course_id} [DELETE]
type APIDeleteUserCourseInput struct {
	required.UserIDField
	Uri APIDeleteUserCourseUri
}
type APIDeleteUserCourseUri struct {
	required.IDField
}

// APIUpdateUserCourseInput /v2/user/course/{course_id} [UPDATE]
type APIUpdateUserCourseInput struct {
	required.UserIDField
	Uri  APIUpdateUserCourseUri
	Body APIUpdateUserCourseBody
}
type APIUpdateUserCourseUri struct {
	required.IDField
}
type APIUpdateUserCourseBody struct {
	optional.NameField
}

// APIGetUserCourseInput /v2/user/course/{course_id} [GET]
type APIGetUserCourseInput struct {
	required.UserIDField
	Uri APIGetUserCourseUri
}
type APIGetUserCourseUri struct {
	required.IDField
}
