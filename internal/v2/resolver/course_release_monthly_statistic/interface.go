package course_release_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_release_monthly_statistic/api_get_cms_statistic_monthly_course_release"
)

type Resolver interface {
	APIGetCMSCourseReleaseStatistic(input *api_get_cms_statistic_monthly_course_release.Input) (output api_get_cms_statistic_monthly_course_release.Output)
	Statistic()
}
