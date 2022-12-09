package api_put_meals

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	mealOptional "github.com/Henry19910227/fitness-go/internal/v2/field/meal/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
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
		foodOptional.IDField
		foodOptional.SourceField
		foodOptional.NameField
		foodOptional.CalorieField
		foodOptional.AmountDescField
		FoodCategory *struct {
			foodCategory.IDField
			foodCategory.TagField
			foodCategory.TitleField
		} `json:"food_category,omitempty"`
	} `json:"food,omitempty"`
}
