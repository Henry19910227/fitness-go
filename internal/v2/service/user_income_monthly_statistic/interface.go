package user_income_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_income_monthly_statistic"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Statistic() (err error)
}
