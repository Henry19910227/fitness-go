package user_unsubscribe_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_unsubscribe_monthly_statistic"

type Resolver interface {
	APIGetCMSUserSubscribeStatistic(input *model.APIGetCMSUserUnsubscribeStatisticInput) (output model.APIGetCMSUserUnsubscribeStatisticOutput)
}
