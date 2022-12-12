package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_create_food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_delete_food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_get_foods"
)

type Resolver interface {
	APICreateFood(input *api_create_food.Input) (output api_create_food.Output)
	APIGetFoods(input *api_get_foods.Input) (output api_get_foods.Output)
	APIDeleteFood(input *api_delete_food.Input) (output api_delete_food.Output)

	APIGetCMSFoods(input *model.APIGetCMSFoodsInput) (output model.APIGetCMSFoodsOutput)
	APICreateCMSFood(input *model.APICreateCMSFoodInput) (output model.APICreateCMSFoodOutput)
	APIUpdateCMSFood(input *model.APIUpdateCMSFoodInput) (output base.Output)
}
