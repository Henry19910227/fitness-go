package workout_set_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_order"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(items []*model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
}
