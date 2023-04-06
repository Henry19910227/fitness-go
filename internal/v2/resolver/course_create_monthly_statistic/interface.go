package course_create_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_create_monthly_statistic/api_get_cms_statistic_monthly_course_create"
)

type Resolver interface {
	APIGetCMSStatisticMonthlyCourseCreate(input *api_get_cms_statistic_monthly_course_create.Input) (output api_get_cms_statistic_monthly_course_create.Output)
	Statistic()
}
