package course

import (
	model "github.com/Henry19910227/fitness-go/internal/model/course"
	"github.com/Henry19910227/fitness-go/internal/model/paging"
)

type Service interface {
	List(input *model.ListParam) (output []*model.Table, page *paging.Output, err error)
	APIGetCMSCourses(input *model.APIGetCMSCoursesInput) interface{}
}
