package admin

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/admin"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
	Update(item *model.Table) (err error)
}
