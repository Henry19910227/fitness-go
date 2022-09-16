package course_release_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_release_monthly_statistics"
}

// APIGetCMSCourseReleaseStatisticOutput /v2/cms/statistic_monthly/course/release [GET]
type APIGetCMSCourseReleaseStatisticOutput struct {
	base.Output
	Data *APIGetCMSCourseReleaseStatisticData `json:"data,omitempty"`
}
type APIGetCMSCourseReleaseStatisticData struct {
	Table
}
