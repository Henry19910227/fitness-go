package favorite_course

import model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_course"

type Repository interface {
	Create(item *model.Table) (err error)
	Delete(input *model.DeleteInput) (err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
