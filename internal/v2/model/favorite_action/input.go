package favorite_action

import "github.com/Henry19910227/fitness-go/internal/v2/field/favorite_aciton/required"

type DeleteInput struct {
	required.UserIDField
	required.ActionIDField
}

// APICreateFavoriteActionInput /v2/favorite/action/{action_id} [POST]
type APICreateFavoriteActionInput struct {
	required.UserIDField
	Uri APICreateFavoriteActionUri
}
type APICreateFavoriteActionUri struct {
	required.ActionIDField
}

// APIDeleteFavoriteActionInput /v2/favorite/action/{action_id} [DELETE]
type APIDeleteFavoriteActionInput struct {
	required.UserIDField
	Uri APIDeleteFavoriteActionUri
}
type APIDeleteFavoriteActionUri struct {
	required.ActionIDField
}
