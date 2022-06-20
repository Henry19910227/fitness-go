package meal

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Create(items []*model.Table) (err error)
	Update(items []*model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
}
