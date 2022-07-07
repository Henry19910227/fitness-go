package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
)

type Service interface {
	List(input *model.ListInput) (outputs []*model.Output, page *paging.Output, err error)
}
