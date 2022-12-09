package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_get_foods"
)

type Resolver interface {
	APIGetFoods(input *api_get_foods.Input) (output api_get_foods.Output)
	APIGetCMSFoods(input *model.APIGetCMSFoodsInput) (output model.APIGetCMSFoodsOutput)
	APICreateCMSFood(input *model.APICreateCMSFoodInput) (output model.APICreateCMSFoodOutput)
	APIUpdateCMSFood(input *model.APIUpdateCMSFoodInput) (output base.Output)
}
