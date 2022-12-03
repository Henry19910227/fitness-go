package banner_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner_order"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Creates(items []*model.Table) (err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Delete(input *model.DeleteInput) (err error)
	DeleteAll() (err error)
}
