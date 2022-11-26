package api_update_trainer_action

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	actionRequired "github.com/Henry19910227/fitness-go/internal/v2/field/action/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
	"github.com/Henry19910227/fitness-go/internal/v2/model/file"
)

// Input /v2/trainer/action/{action_id} [PATCH] 修改教練動作 API
type Input struct {
	userRequired.UserIDField
	Cover *file.Input
	Video *file.Input
	Form  Form
	Uri   Uri
}
type Form struct {
	actionOptional.NameField
	actionOptional.CategoryField
	actionOptional.BodyField
	actionOptional.EquipmentField
	actionOptional.IntroField
}
type Uri struct {
	actionRequired.IDField
}