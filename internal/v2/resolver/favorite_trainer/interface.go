package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer/api_create_favorite_trainer"
)

type Resolver interface {
	APICreateFavoriteTrainer(input *api_create_favorite_trainer.Input) (output api_create_favorite_trainer.Output)
}
