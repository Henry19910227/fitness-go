package api_create_trainer_action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/field/action/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/action [POST] 新增教練動作 API
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	optional.IDField
}
