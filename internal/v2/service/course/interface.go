package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (id int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, page *paging.Output, err error)
	Updates(items []*model.Table) (err error)
	Update(item *model.Table) (err error)
}
