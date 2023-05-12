package subscribe_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	CreateOrUpdate(item *model.Table) (id *int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
	Create(item *model.Table) (id int64, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
