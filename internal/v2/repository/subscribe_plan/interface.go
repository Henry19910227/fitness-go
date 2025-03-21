package subscribe_plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
