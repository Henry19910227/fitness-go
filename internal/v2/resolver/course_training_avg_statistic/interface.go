package course_training_avg_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic"

type Resolver interface {
	APIGetCMSCourseTrainingAvgStatistic(input *model.APIGetCMSCourseTrainingAvgStatisticInput) (output model.APIGetCMSCourseTrainingAvgStatisticOutput)
}
