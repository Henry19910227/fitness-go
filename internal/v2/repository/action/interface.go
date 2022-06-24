package action

import model "github.com/Henry19910227/fitness-go/internal/v2/model/action"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	Create(item *model.Table) (id int64, err error)
}
