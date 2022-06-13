package course

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course"

type Resolver interface {
	APIGetCMSCourses(input *model.APIGetCMSCoursesInput) interface{}
	APIGetCMSCourse(input *model.APIGetCMSCourseInput) interface{}
}
