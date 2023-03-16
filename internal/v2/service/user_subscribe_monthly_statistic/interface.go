package user_subscribe_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Statistic(input *model.StatisticInput) (err error)
}
