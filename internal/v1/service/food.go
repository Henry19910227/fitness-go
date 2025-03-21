package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
)

type food struct {
	Base
	foodRepo         repository.Food
	foodCategoryRepo repository.FoodCategory
	calorieTool tool.Calorie
	errHandler  errcode.Handler
}

func NewFood(foodRepo repository.Food, foodCategoryRepo repository.FoodCategory, calorieTool tool.Calorie, errHandler errcode.Handler) Food {
	return &food{foodRepo: foodRepo, foodCategoryRepo: foodCategoryRepo, calorieTool: calorieTool, errHandler: errHandler}
}

func (f *food) CreateFood(c *gin.Context, param *dto.CreateFoodParam) (*dto.Food, errcode.Error) {
	//查找category
	category, err := f.foodCategoryRepo.FindFoodCategory(param.FoodCategoryID)
	if err != nil {
		return nil, f.errHandler.Set(c, "food category repo", err)
	}
	//創建food
	foodID, err := f.foodRepo.CreateFood(&model.CreateFoodParam{
		UserID:         util.PointerInt64(param.UserID),
		FoodCategoryID: param.FoodCategoryID,
		Source:         param.Source,
		Name:           param.Name,
		Calorie:        f.calorieTool.FoodCalorie(global.FoodCategoryTag(category.Tag)),
		AmountDesc:     param.AmountDesc,
	})
	if err != nil {
		return nil, f.errHandler.Set(c, "food repo", err)
	}
	//查找food
	preloads := make([]*model.Preload, 0)
	preloads = append(preloads, &model.Preload{Field: "FoodCategory"})
	data, err := f.foodRepo.FindFood(foodID, preloads)
	if err != nil {
		return nil, f.errHandler.Set(c, "food repo", err)
	}
	//parser
	var food dto.Food
	if err := util.Parser(data, &food); err != nil {
		return nil, f.errHandler.Set(c, "parser error", err)
	}
	return &food, nil
}

func (f *food) GetFoods(c *gin.Context, userID int64, tag global.FoodCategoryTag) ([]*dto.Food, errcode.Error) {
	input := model.FindFoodsParam{}
	input.IsDeleted = util.PointerInt(0)
	input.UserID = util.PointerInt64(userID)
	input.Tag = util.PointerInt(int(tag))
	input.Preloads = []*model.Preload{{Field: "FoodCategory"}}
	datas, err := f.foodRepo.FindFoods(&input)
	if err != nil {
		return nil, f.errHandler.Set(c, "food repo", err)
	}
	foods := make([]*dto.Food, 0)
	if err := util.Parser(datas, &foods); err != nil {
		return nil, f.errHandler.Set(c, "foods parser error", err)
	}
	return foods, nil
}

func (f *food) GetRecentFood(c *gin.Context, userID int64) ([]*dto.RecentFood, errcode.Error) {
	param := model.FindRecentFoodsParam{}
	param.UserID = util.PointerInt64(userID)
	param.IsDeleted = util.PointerInt(0)
	param.Preloads = []*model.Preload{{Field: "FoodCategory"}, {Field: "Meal"}}
	datas, err := f.foodRepo.FindRecentFoods(&param)
	if err != nil {
		return nil, f.errHandler.Set(c, "food repo", err)
	}
	foods := make([]*dto.RecentFood, 0)
	if err := util.Parser(datas, &foods); err != nil {
		return nil, f.errHandler.Set(c, "foods parser error", err)
	}
	return foods, nil
}

func (f *food) DeleteFood(c *gin.Context, foodID int64, userID int64) errcode.Error {
	ownerID, err := f.foodRepo.FindFoodOwner(foodID)
	if err != nil {
		return f.errHandler.Set(c, "foods parser error", err)
	}
	if ownerID != userID {
		return f.errHandler.PermissionDenied()
	}
	if err := f.foodRepo.UpdateFood(&model.UpdateFoodParam{
		FoodID:    foodID,
		IsDeleted: util.PointerInt(1),
	}); err != nil {
		return f.errHandler.Set(c, "food repo", err)
	}
	return nil
}
