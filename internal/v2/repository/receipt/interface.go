package receipt

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
