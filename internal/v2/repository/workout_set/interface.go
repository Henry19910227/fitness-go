package workout_set

import model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"

type Repository interface {
	List(input *model.ListInput) (output []*model.Table, amount int64, err error)
}
