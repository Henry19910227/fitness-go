package api_update_trainer_plan

import (
	planOptional "github.com/Henry19910227/fitness-go/internal/v2/field/plan/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/trainer/plan/{plan_id} [PATCH]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	planOptional.IDField
	planOptional.NameField
	planOptional.WorkoutCountField
	planOptional.CreateAtField
	planOptional.UpdateAtField
}
