package api_update_trainer_plan

import (
	planRequired "github.com/Henry19910227/fitness-go/internal/v2/field/plan/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/trainer/plan/{plan_id} [PATCH]
type Input struct {
	userRequired.UserIDField
	Uri  Uri
	Body Body
}
type Uri struct {
	planRequired.IDField
}
type Body struct {
	planRequired.NameField
}
