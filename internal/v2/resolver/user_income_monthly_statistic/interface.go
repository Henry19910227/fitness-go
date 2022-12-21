package user_income_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_income_monthly_statistic/api_get_trainer_income_monthly_statistic"
)

type Resolver interface {
	APIGetTrainerIncomeMonthlyStatistic(input *api_get_trainer_income_monthly_statistic.Input) (output api_get_trainer_income_monthly_statistic.Output)
}
