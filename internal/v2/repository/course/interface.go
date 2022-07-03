package course

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
)

type Repository interface {
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, amount int64, err error)
	FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, amount int64, err error)
	Updates(items []*model.Table) (err error)
	Update(item *model.Table) (err error)
}
