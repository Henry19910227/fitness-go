package course_category_training_monthly_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Statistic(input *model.StatisticInput) (err error)
}
