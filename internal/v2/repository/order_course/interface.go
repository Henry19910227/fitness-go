package order_course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (err error)
}
