package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
)

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (output []*model.Output, amount int64, err error)
}
