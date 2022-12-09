package food

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Create(item *model.Table) (output *model.Output, err error)
	Update(item *model.Table) (err error)
}
