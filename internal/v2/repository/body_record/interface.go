package body_record

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
)

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	Create(item *model.Table) (id int64, err error)
}
