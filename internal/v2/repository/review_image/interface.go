package review_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	Create(items []*model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
}
