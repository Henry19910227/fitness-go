package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
)

type Resolver interface {
	APIGetFoods(input *model.APIGetFoodsInput) (output model.APIGetFoodsOutput)
	APIGetCMSFoods() (output model.APIGetCMSFoodsOutput)
	APICreateCMSFood(input *model.APICreateCMSFoodInput) (output model.APICreateCMSFoodOutput)
	APIUpdateCMSFood(input *model.APIUpdateCMSFoodInput) (output base.Output)
}
