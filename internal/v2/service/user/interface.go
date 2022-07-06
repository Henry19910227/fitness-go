package user

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
)

type Service interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
}
