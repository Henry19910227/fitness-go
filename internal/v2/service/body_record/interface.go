package body_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	Create(item *model.Table) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	LatestList(input *model.LatestListInput) (outputs []*model.Output, err error)
	Update(item *model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
}
