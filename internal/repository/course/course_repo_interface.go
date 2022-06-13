package course

import (
	model "github.com/Henry19910227/fitness-go/internal/model/course"
)

type Repository interface {
	List(input *model.ListParam) (output []*model.Table, amount int64, err error)
}
