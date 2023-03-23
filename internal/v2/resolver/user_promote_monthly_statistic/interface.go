package user_promote_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_promote_monthly_statistic/api_get_cms_statistic_monthly_user_promote"
)

type Resolver interface {
	APIGetCMSUserPromoteStatistic(input *api_get_cms_statistic_monthly_user_promote.Input) (output api_get_cms_statistic_monthly_user_promote.Output)
	Statistic()
}
