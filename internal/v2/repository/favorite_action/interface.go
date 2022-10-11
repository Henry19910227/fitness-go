package favorite_action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"
)

type Repository interface {
	Create(item *model.Table) (err error)
}
