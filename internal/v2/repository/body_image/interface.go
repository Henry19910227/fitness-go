package body_image

import model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"

type Repository interface {
	Create(item *model.Table) (id int64, err error)
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
