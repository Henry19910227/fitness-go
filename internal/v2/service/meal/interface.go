package meal

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(items []*model.Table) (err error)
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
	Delete(input *model.DeleteInput) (err error)
}
