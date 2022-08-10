package order_subscribe_plan

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (err error)
	Update(item *model.Table) (err error)
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
}
