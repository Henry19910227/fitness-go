package user

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
	Create(item *model.Table) (id int64, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
