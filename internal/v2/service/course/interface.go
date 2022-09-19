package course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/entity/course"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *course.Table) (id int64, err error)
	Delete(input *model.DeleteInput) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, page *paging.Output, err error)
	ProgressList(input *model.ProgressListInput) (outputs []*model.Output, page *paging.Output, err error)
	ChargeList(input *model.ChargeListInput) (outputs []*model.Output, page *paging.Output, err error)
	Updates(items []*course.Table) (err error)
	Update(item *course.Table) (err error)
}
