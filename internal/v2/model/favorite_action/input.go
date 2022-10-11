package favorite_action

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_aciton/required"

// APICreateFavoriteActionInput /v2/favorite/action/{action_id} [POST]
type APICreateFavoriteActionInput struct {
	required.UserIDField
	Uri APICreateFavoriteActionUri
}
type APICreateFavoriteActionUri struct {
	required.ActionIDField
}
