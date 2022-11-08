package api_delete_favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/favorite_trainer/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/favorite/trainer/{user_id} [DELETE]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.TrainerIDField
}
