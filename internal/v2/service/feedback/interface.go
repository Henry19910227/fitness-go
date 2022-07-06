package feedback

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Create(item *model.Table) (id int64, err error)
}
