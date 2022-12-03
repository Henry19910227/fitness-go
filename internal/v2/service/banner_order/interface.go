package banner_order

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner_order"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Creates(items []*model.Table) (err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Delete(input *model.DeleteInput) (err error)
	DeleteAll() (err error)
}
