package course_category_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic/api_get_cms_statistic_monthly_course_category_training"
)

type Resolver interface {
	APIGetCMSCategoryTrainingStatistic(input *api_get_cms_statistic_monthly_course_category_training.Input) (output api_get_cms_statistic_monthly_course_category_training.Output)
	Statistic()
}
