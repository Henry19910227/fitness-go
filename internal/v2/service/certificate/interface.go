package certificate

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (id int64, err error)
	Creates(items []*model.Table) (err error)
	Updates(items []*model.Table) (err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Deletes(inputs []*model.DeleteInput) (err error)
}
