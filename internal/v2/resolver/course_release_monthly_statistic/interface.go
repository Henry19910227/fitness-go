package course_release_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/course_release_monthly_statistic"

type Resolver interface {
	APIGetCMSCourseReleaseStatistic(input *model.APIGetCMSCourseReleaseStatisticInput) (output model.APIGetCMSCourseReleaseStatisticOutput)
}
