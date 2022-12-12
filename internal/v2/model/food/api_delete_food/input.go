package api_delete_food

import (
	foodRequired "github.com/Henry19910227/fitness-go/internal/v2/field/food/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/food/{food_id} [DELETE]
type Input struct {
	userRequired.UserIDField
	Uri Uri
}
type Uri struct {
	foodRequired.FoodIDField
}
