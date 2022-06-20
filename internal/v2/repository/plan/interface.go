package plan

import model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"

type Repository interface {
	List(input *model.ListInput) (output []*model.Output, amount int64, err error)
}
