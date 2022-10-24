package review

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Update(item *model.Table) (err error)
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
	Delete(input *model.DeleteInput) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Create(item *model.Table) (id int64, err error)
}
