package user_register_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_register_monthly_statistic"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Statistic(input *model.StatisticInput) (err error)
}
