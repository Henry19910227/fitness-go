package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Create(item *model.Table) (err error)
	CreateOrUpdate(item *model.Table) (err error)
	Update(item *model.Table) (err error)
	Updates(items []*model.Table) (err error)
}
