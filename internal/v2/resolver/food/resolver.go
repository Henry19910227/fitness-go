package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	foodCalorie "github.com/Henry19910227/fitness-go/internal/pkg/tool/food_calorie"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_create_food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_get_foods"
	categoryModel "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	foodCategoryService "github.com/Henry19910227/fitness-go/internal/v2/service/food_category"
)

type resolver struct {
	foodService         foodService.Service
	foodCategoryService foodCategoryService.Service
	calorieTool         foodCalorie.Tool
}

func New(foodService foodService.Service, foodCategoryService foodCategoryService.Service, calorieTool foodCalorie.Tool) Resolver {
	return &resolver{foodService: foodService, foodCategoryService: foodCategoryService, calorieTool: calorieTool}
}

func (r *resolver) APICreateFood(input *api_create_food.Input) (output api_create_food.Output) {
	// 查詢 food category
	findInput := categoryModel.FindInput{}
	findInput.ID = util.PointerInt64(input.Body.FoodCategoryID)
	foodCategoryOutput, err := r.foodCategoryService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 計算卡路里
	cal := r.calorieTool.Calorie(util.OnNilJustReturnInt(foodCategoryOutput.Tag, 0))
	// 創建 Food
	table := model.Table{}
	table.Source = util.PointerInt(model.Custom)
	table.Calorie = util.PointerInt(cal)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	_, err = r.foodService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetFoods(input *api_get_foods.Input) (output api_get_foods.Output) {
	// parser input
	listInput := model.ListInput{}
	listInput.Name = input.Query.Name
	listInput.Status = util.PointerInt(1)
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	wheres := make([]*whereModel.Where, 0)
	joins := make([]*joinModel.Join, 0)
	wheres = append(wheres, &whereModel.Where{Query: "(foods.user_id = ? OR foods.user_id IS NULL)", Args: []interface{}{input.UserID}})
	if input.Query.Tag != nil {
		joins = append(joins, &joinModel.Join{Query: "INNER JOIN food_categories ON foods.food_category_id = food_categories.id"})
		wheres = append(wheres, &whereModel.Where{Query: "food_categories.tag = ?", Args: []interface{}{*input.Query.Tag}})
	}
	listInput.Wheres = wheres
	listInput.Joins = joins
	// 調用 service
	result, _, err := r.foodService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_foods.Data{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
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
