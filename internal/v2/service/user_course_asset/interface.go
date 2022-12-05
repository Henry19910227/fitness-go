package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
	Create(item *model.Table) (id int64, err error)
	Creates(items []*model.Table) (err error)
}
