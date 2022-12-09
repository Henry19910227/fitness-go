package api_get_foods

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/foods [GET]
type Input struct {
	userRequired.UserIDField
	Query Query
}
type Query struct {
	foodCategoryOptional.TagField
	foodOptional.UserIDField
	foodOptional.NameField
}
