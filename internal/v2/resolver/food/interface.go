package food

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
)

type Resolver interface {
	APIGetFoods(input *model.APIGetFoodsInput) (output model.APIGetFoodsOutput)
}
