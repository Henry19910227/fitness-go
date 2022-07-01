package food_category

import model "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"

type Resolver interface {
	APIGetCMSFoodCategories() (output model.APIGetCMSFoodCategoriesOutput)
}
