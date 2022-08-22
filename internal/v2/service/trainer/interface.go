package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
)

type Service interface {
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, page *paging.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
}
