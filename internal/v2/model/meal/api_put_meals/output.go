package api_put_meals

import (
	mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
)

type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data struct {
	mealOptional.IDField
	mealOptional.TypeField
	mealOptional.AmountField
	mealOptional.CreateAtField
	Food *struct {
		food.IDField
		food.SourceField
		food.NameField
		food.CalorieField
		food.AmountDescField
		FoodCategory *struct {
			foodCategory.IDField
			foodCategory.TagField
			foodCategory.TitleField
		} `json:"food_category,omitempty"`
	} `json:"food,omitempty"`
}
