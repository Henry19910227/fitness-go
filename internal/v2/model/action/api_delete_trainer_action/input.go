package api_delete_trainer_action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/trainer/action/{action_id} [DELETE]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	required.ActionIDField
}
