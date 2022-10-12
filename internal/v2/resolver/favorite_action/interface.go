package favorite_action

import model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_action"

type Resolver interface {
	APICreateFavoriteAction(input *model.APICreateFavoriteActionInput) (output model.APICreateFavoriteActionOutput)
	APIDeleteFavoriteAction(input *model.APIDeleteFavoriteActionInput) (output model.APIDeleteFavoriteActionOutput)
}
