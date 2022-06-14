package plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
)

type Service interface {
	List(input *model.ListInput) (output []*model.Table, page *paging.Output, err error)
}
