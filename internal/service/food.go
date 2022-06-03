package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/gin-gonic/gin"
)

type food struct {
	Base
	foodRepo repository.Food
	foodCategoryRepo repository.FoodCategory
	calorieTool tool.Calorie
	errHandler errcode.Handler
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
		UserID: util.PointerInt64(param.UserID),
		FoodCategoryID: param.FoodCategoryID,
		Source: param.Source,
		Name: param.Name,
		Calorie: f.calorieTool.FoodCalorie(global.FoodCategoryTag(category.Tag)),
		AmountDesc: param.AmountDesc,
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
