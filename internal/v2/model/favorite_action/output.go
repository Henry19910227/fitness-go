package favorite_action

import "github.com/Henry19910227/fitness-go/internal/v2/model/base"

// APICreateFavoriteActionOutput /v2/favorite/action/{action_id} [POST]
type APICreateFavoriteActionOutput struct {
	base.Output
}

// APIDeleteFavoriteActionOutput /v2/favorite/action/{action_id} [DELETE]
type APIDeleteFavoriteActionOutput struct {
	base.Output
}
