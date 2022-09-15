package user_subscribe_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic"

type Resolver interface {
	APIGetCMSUserSubscribeStatistic(input *model.APIGetCMSUserSubscribeStatisticInput) (output model.APIGetCMSUserSubscribeStatisticOutput)
}
