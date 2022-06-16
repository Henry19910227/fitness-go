package food

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
