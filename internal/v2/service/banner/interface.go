package banner

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Service interface {
	Create(item *model.Table) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
	Delete(input *model.DeleteInput) (err error)
	Find(input *model.FindInput) (output *model.Output, err error)
}
