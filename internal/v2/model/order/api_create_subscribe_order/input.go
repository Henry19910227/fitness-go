package api_create_subscribe_order

import (
	subscribePlanRequired "github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_plan/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/subscribe_order [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	subscribePlanRequired.SubscribePlanIDField
}
