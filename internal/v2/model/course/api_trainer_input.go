package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/course/required"
)

// APIGetTrainerCoursesInput /v2/trainer/courses [GET]
type APIGetTrainerCoursesInput struct {
	required.UserIDField
	Query APIGetTrainerCoursesQuery
}
type APIGetTrainerCoursesQuery struct {
	optional.CourseStatusField
	PagingInput
}

// APICreateTrainerCourseInput /v2/trainer/course [POST]
type APICreateTrainerCourseInput struct {
	required.UserIDField
	Body APICreateTrainerCourseBody
}
type APICreateTrainerCourseBody struct {
	required.NameField
	required.CategoryField
	required.LevelField
	required.ScheduleTypeField
}
