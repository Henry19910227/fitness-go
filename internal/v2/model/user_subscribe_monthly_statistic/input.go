package user_subscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/user_subscribe_monthly_statistic/required"
)

type StatisticInput struct {
	required.YearField
	required.MonthField
}
