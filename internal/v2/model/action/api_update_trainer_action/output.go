package api_update_trainer_action

import (
	actionOptional "github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/action/{action_id} [PATCH] 修改教練動作 API
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	actionOptional.IDField
	actionOptional.NameField
	actionOptional.TypeField
	actionOptional.CategoryField
	actionOptional.BodyField
	actionOptional.EquipmentField
	actionOptional.IntroField
	actionOptional.CoverField
	actionOptional.VideoField
	actionOptional.StatusField
	actionOptional.CreateAtField
	actionOptional.UpdateAtField
}
