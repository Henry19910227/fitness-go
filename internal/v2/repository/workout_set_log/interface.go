package workout_set_log

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(items []*model.Table) (ids []int64, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
