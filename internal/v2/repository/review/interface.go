package review

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Update(item *model.Table) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Delete(input *model.DeleteInput) (err error)
	Create(item *model.Table) (id int64, err error)
}
