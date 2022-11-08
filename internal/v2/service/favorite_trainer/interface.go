package favorite_trainer

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	Create(item *model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
