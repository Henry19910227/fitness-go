package review

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
)

type Service interface {
	Update(item *model.Table) (err error)
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
}
