package user_unsubscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_unsubscribe_monthly_statistic/api_get_cms_statistic_monthly_user_unsubscribe"
)

type Resolver interface {
	APIGetCMSStatisticMonthlyUserUnsubscribe(input *api_get_cms_statistic_monthly_user_unsubscribe.Input) (output api_get_cms_statistic_monthly_user_unsubscribe.Output)
}
