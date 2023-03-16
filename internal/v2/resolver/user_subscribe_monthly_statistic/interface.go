package user_subscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic/api_get_cms_user_subscribe_statistic"
)

type Resolver interface {
	APIGetCMSUserSubscribeStatistic(input *api_get_cms_user_subscribe_statistic.Input) (output api_get_cms_user_subscribe_statistic.Output)
	Statistic()
}
