package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic/api_get_cms_statistic_monthly_course_training_avg"
)

type Resolver interface {
	APIGetCMSCourseTrainingAvgStatistic(input *api_get_cms_statistic_monthly_course_training_avg.Input) (output api_get_cms_statistic_monthly_course_training_avg.Output)
	Statistic()
}
