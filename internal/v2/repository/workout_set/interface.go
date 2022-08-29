package workout_set

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	Create(items []*model.Table) (ids []int64, err error)
	List(input *model.ListInput) (output []*model.Output, amount int64, err error)
	Delete(input *model.DeleteInput) (err error)
}
