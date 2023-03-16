package user_subscribe_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Statistic(input *model.StatisticInput) (err error)
}
