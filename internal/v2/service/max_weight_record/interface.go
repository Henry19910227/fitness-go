package max_weight_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/max_weight_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Create(item *model.Table) (id int64, err error)
	Update(item *model.Table) (err error)
	CreateOrUpdate(item *model.Table) (id *int64, err error)
}
