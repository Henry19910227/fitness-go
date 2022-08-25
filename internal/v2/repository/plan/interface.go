package plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (id int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
	List(input *model.ListInput) (output []*model.Output, amount int64, err error)
}
