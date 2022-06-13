package course

import (
	model "github.com/Henry19910227/fitness-go/internal/model/course"
	"github.com/Henry19910227/fitness-go/internal/model/paging"
)

type Service interface {
	Find(input *model.FindInput) (output *model.Table, err error)
	List(input *model.ListInput) (output []*model.Table, page *paging.Output, err error)
}
