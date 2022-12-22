package user_subscribe_info

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Create(item *model.Table) (err error)
	CreateOrUpdate(item *model.Table) (err error)
	Update(item *model.Table) (err error)
	Updates(items []*model.Table) (err error)
}
