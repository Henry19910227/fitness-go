package food_category

import model "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
}
