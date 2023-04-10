package course_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_monthly_statistic/api_get_cms_statistic_monthly_course_training"
)

type Resolver interface {
	APIGetCMSCourseTrainingStatistic(input *api_get_cms_statistic_monthly_course_training.Input) (output api_get_cms_statistic_monthly_course_training.Output)
	Statistic()
}
