package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
)

type Service interface {
	List(input *model.ListInput) (output []*model.Output, page *paging.Output, err error)
}
