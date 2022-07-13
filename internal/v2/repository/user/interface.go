package user

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	Update(item *model.Table) (err error)
	Create(item *model.Table) (id int64, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
