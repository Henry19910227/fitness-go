package api_get_subscribe_plans

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/sale_items [GET]
type Input struct {
	userRequired.UserIDField
}
