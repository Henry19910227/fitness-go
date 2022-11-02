package bank_account

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	Create(item *model.Table) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Update(item *model.Table) (err error)
}
