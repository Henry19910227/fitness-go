package user_promote_monthly_statistic

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_promote_monthly_statistic"

type Resolver interface {
	APIGetCMSUserPromoteStatistic(input *model.APIGetCMSUserPromoteStatisticInput) (output model.APIGetCMSUserPromoteStatisticOutput)
}
