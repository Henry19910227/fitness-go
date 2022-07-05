package feedback_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(items []*model.Table) (err error)
}
