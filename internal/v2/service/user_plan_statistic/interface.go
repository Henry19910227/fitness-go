package user_plan_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_plan_statistic"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Statistic(input *model.Statistic) (err error)
}
