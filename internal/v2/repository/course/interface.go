package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
)

type Repository interface {
	Find(input *model.FindInput) (output *model.Table, err error)
	List(input *model.ListInput) (output []*model.Table, amount int64, err error)
}
