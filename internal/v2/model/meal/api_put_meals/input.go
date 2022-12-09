package api_put_meals

import (
	mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"
	userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"
)

// Input /v2/diet/{diet_id}/meals [PUT] 修改並覆蓋餐食 API
type Input struct {
	userRequired.UserIDField
	Uri  Uri
	Body Body
}
type Uri struct {
	mealOptional.DietIDField
}
type Body []*struct {
	mealOptional.FoodIDField
	mealOptional.TypeField
	mealOptional.AmountField
}
