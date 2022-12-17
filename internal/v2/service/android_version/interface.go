package android_version

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/android_version"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
}
