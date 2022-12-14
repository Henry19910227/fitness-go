package food_category

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food_category/api_get_food_categories"
)

type Resolver interface {
	APIGetFoodCategories(input *api_get_food_categories.Input) (output api_get_food_categories.Output)
	APIGetCMSFoodCategories() (output model.APIGetCMSFoodCategoriesOutput)
}
