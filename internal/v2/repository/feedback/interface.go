package feedback

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (id int64, err error)
}
