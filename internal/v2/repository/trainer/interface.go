package trainer

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Update(item *model.Table) (err error)
}
