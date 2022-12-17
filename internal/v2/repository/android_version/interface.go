package android_version

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/android_version"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
