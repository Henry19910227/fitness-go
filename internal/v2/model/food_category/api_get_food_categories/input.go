package api_get_food_categories

import userRequired "github.com/Henry19910227/fitness-go/internal/v2/field/user/required"

// Input /v2/food_categories [GET] 獲取食物分類
type Input struct {
	userRequired.UserIDField
}
