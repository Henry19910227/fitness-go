package api_get_foods

import (
	foodOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food/optional"
	foodCategoryOptional "github.com/Henry19910227/fitness-go/internal/v2/field/food_category/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

// Output /v2/foods [GET]
type Output struct {
	base.Output
	Data *Data `json:"data,omitempty"`
}
type Data []*struct {
	foodOptional.IDField
	foodOptional.UserIDField
	foodOptional.SourceField
	foodOptional.NameField
	foodOptional.CalorieField
	foodOptional.AmountDescField
	foodOptional.CreateAtField
	foodOptional.UpdateAtField
	FoodCategory *struct {
		foodCategoryOptional.IDField
		foodCategoryOptional.TagField
		foodCategoryOptional.TitleField
	} `json:"food_category,omitempty"`
}
