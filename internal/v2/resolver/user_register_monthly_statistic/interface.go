package user_register_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_register_monthly_statistic/api_get_cms_statistic_monthly_user_register"
)

type Resolver interface {
	APIGetCMSUserRegisterStatistic(input *api_get_cms_statistic_monthly_user_register.Input) (output api_get_cms_statistic_monthly_user_register.Output)
	Statistic()
}
