package course_training_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_monthly_statistic"

type Resolver interface {
	APIGetCMSCourseTrainingStatistic(input *model.APIGetCMSCourseTrainingStatisticInput) (output model.APIGetCMSCourseTrainingStatisticOutput)
}