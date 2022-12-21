package user_income_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_income_monthly_statistic"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
