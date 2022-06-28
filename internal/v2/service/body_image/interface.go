package body_image

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	Create(item *model.Table) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
