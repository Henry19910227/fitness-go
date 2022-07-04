package trainer

import model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, amount int64, err error)
	Update(item *model.Table) (err error)
}
