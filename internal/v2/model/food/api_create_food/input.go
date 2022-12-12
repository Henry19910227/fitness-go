package api_create_food

import (
	foodRequired "github.com/Henry19910227/fitness-go/internal/v2/field/food/required"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/food [POST]
type Input struct {
	userRequired.UserIDField
	Body Body
}
type Body struct {
	foodRequired.FoodCategoryIDField
	foodRequired.NameField
	foodRequired.AmountDescField
}
