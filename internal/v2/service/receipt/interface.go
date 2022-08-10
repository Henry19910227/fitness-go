package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
