package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(items []*model.Table) (ids []int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Delete(input *model.DeleteInput) (err error)
}
