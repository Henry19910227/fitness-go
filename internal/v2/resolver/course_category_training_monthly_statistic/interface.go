package course_category_training_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic"

type Resolver interface {
	APIGetCMSCategoryTrainingStatistic(input *model.APIGetCMSCategoryTrainingStatisticInput) (output model.APIGetCMSCategoryTrainingStatisticOutput)
}
