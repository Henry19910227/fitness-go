package course_training_monthly_statistic

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

type Output struct {
	Table
}

func (Output) TableName() string {
	return "course_training_monthly_statistics"
}

// APIGetCMSCourseTrainingStatisticOutput /v2/cms/statistic_monthly/course/training [GET]
type APIGetCMSCourseTrainingStatisticOutput struct {
	base.Output
	Data *APIGetCMSCourseTrainingStatisticData `json:"data,omitempty"`
}
type APIGetCMSCourseTrainingStatisticData struct {
	Table
}
