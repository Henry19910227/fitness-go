package course_create_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_create_monthly_statistics"
}

// APIGetCMSCourseCreateStatisticOutput /v2/cms/statistic_monthly/course/create [GET]
type APIGetCMSCourseCreateStatisticOutput struct {
	base.Output
	Data *APIGetCMSCourseCreateStatisticData `json:"data,omitempty"`
}
type APIGetCMSCourseCreateStatisticData struct {
	Table
}
