package card

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/card"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (err error)
}
