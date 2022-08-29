package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Create(item *model.Table) (id int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, amount int64, err error)
	ProgressList(input *model.ProgressListInput) (outputs []*model.Output, amount int64, err error)
	ChargeList(input *model.ChargeListInput) (outputs []*model.Output, amount int64, err error)
	Updates(items []*model.Table) (err error)
	Update(item *model.Table) (err error)
}
