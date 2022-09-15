package user_register_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_register_monthly_statistic"

type Resolver interface {
	APIGetCMSUserRegisterStatistic(input *model.APIGetCMSUserRegisterStatisticInput) (output model.APIGetCMSUserRegisterStatisticOutput)
}
