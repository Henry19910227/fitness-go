package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	foodCalorie "github.com/Henry19910227/fitness-go/internal/pkg/tool/food_calorie"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	categoryModel "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	foodCategoryService "github.com/Henry19910227/fitness-go/internal/v2/service/food_category"
)

type resolver struct {
	foodService foodService.Service
	foodCategoryService foodCategoryService.Service
	calorieTool foodCalorie.Tool
}

func New(foodService foodService.Service, foodCategoryService foodCategoryService.Service, calorieTool foodCalorie.Tool) Resolver {
	return &resolver{foodService: foodService, foodCategoryService: foodCategoryService, calorieTool: calorieTool}
}

func (r *resolver) APIGetFoods(input *model.APIGetFoodsInput) (output model.APIGetFoodsOutput) {
	// parser input
	param := model.ListInput{}
	param.Status = util.PointerInt(1)
	param.OrderField = "create_at"
	param.OrderType = order_by.DESC
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

func (r *resolver) APIGetCMSFoods(input *model.APIGetCMSFoodsInput) (output model.APIGetCMSFoodsOutput) {
	// parser input
	param := model.ListInput{}
	param.Source = util.PointerInt(1)
	param.Size = input.Form.Size
	param.Page = input.Form.Page
	param.OrderField = "create_at"
	param.OrderType = order_by.ASC
	param.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	// 調用 service
	result, page, err := r.foodService.List(&param)
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
	output.Paging = page
	return output
}

func (r *resolver) APICreateCMSFood(input *model.APICreateCMSFoodInput) (output model.APICreateCMSFoodOutput) {
	//查找 food category並計算卡路里
	findInput := categoryModel.FindInput{}
	findInput.ID = util.PointerInt64(input.Body.FoodCategoryID)
	foodCategoryOutput, err := r.foodCategoryService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	cal := r.calorieTool.Calorie(util.OnNilJustReturnInt(foodCategoryOutput.Tag, 0))
	//parser input
	table := model.Table{}
	table.Source = util.PointerInt(model.System)
	table.Calorie = util.PointerInt(cal)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//執行創建
	result, err := r.foodService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser output
	data := model.APICreateCMSFoodData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateCMSFood(input *model.APIUpdateCMSFoodInput) (output base.Output) {
	//parser input
	table := model.Table{}
	table.ID = util.PointerInt64(input.Uri.ID)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//驗證是否修改的食物為系統食物
	findParam := model.FindInput{}
	findParam.ID = util.PointerInt64(input.Uri.ID)
	findOutput, err := r.foodService.Find(&findParam)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnInt(findOutput.Source, 0) != model.System {
		output.Set(code.PermissionDenied, err.Error())
		return output
	}
	//執行更新
	if err := r.foodService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}
