package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
)

type resolver struct {
	foodService foodService.Service
}

func New(foodService foodService.Service) Resolver {
	return &resolver{foodService: foodService}
}

func (r *resolver) APIGetFoods(input *model.APIGetFoodsInput) (output model.APIGetFoodsOutput) {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	// 調用 repo
	result, _, err := r.foodService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetFoodsData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}
