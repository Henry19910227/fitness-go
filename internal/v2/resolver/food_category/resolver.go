package food_category

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food_category/api_get_food_categories"
	foodCategoryService "github.com/Henry19910227/fitness-go/internal/v2/service/food_category"
)

type resolver struct {
	foodService foodCategoryService.Service
}

func New(foodService foodCategoryService.Service) Resolver {
	return &resolver{foodService: foodService}
}

func (r *resolver) APIGetFoodCategories(input *api_get_food_categories.Input) (output api_get_food_categories.Output) {
	listInput := model.ListInput{}
	categoryOutputs, page, err := r.foodService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := api_get_food_categories.Data{}
	if err := util.Parser(categoryOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	output.Paging = page
	return output
}

func (r *resolver) APIGetCMSFoodCategories() (output model.APIGetCMSFoodCategoriesOutput) {
	param := model.ListInput{}
	// 調用 service
	result, _, err := r.foodService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSFoodCategoriesData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	return output
}
