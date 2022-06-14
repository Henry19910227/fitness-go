package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
)

type foodCategory struct {
	Base
	foodCategoryRepo repository.FoodCategory
	errHandler       errcode.Handler
}

func NewFoodCategory(foodCategoryRepo repository.FoodCategory, errHandler errcode.Handler) FoodCategory {
	return &foodCategory{foodCategoryRepo: foodCategoryRepo, errHandler: errHandler}
}

func (f *foodCategory) GetFoodCategories(c *gin.Context) ([]*dto.FoodCategory, errcode.Error) {
	datas, err := f.foodCategoryRepo.FindFoodCategories()
	if err != nil {
		return nil, f.errHandler.Set(c, "food_category_repo", err)
	}
	categories := make([]*dto.FoodCategory, 0)
	if err := util.Parser(datas, &categories); err != nil {
		return nil, f.errHandler.Set(c, "parser error", err)
	}
	return categories, nil
}
