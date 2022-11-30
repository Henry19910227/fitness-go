package api_update_cms_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/required"
)

// Input /v2/cms/trainer/{user_id} [PATCH]
type Input struct {
	Uri  Uri
	Body Body
}
type Uri struct {
	required.UserIDField
}
type Body struct {
	optional.TrainerStatusField
	optional.TrainerLevelField
}
