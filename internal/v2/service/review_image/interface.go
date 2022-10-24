package review_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	Delete(input *model.DeleteInput) (err error)
	Create(items []*model.Table) (err error)
}
