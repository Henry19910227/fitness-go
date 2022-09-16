package course_create_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course_create_monthly_statistic"

type Resolver interface {
	APIGetCMSCourseCreateStatistic(input *model.APIGetCMSCourseCreateStatisticInput) (output model.APIGetCMSCourseCreateStatisticOutput)
}
