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
	param.Status = util.PointerInt(1)
	if err := util.Parser(input, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	// 調用 service
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

func (r *resolver) APIGetCMSFoods() (output model.APIGetCMSFoodsOutput) {
	// parser input
	param := model.ListInput{}
	param.Source = util.PointerInt(1)
	param.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	// 調用 service
	result, _, err := r.foodService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSFoodsData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}

func (r *resolver) APICreateCMSFood(input *model.APICreateCMSFoodInput) (output model.APICreateCMSFoodOutput) {
	table := model.Table{}
	table.Source = util.PointerInt(1)
	table.Calorie = util.PointerInt(0)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	result, err := r.foodService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateCMSFoodData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
