package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/util"
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
	for _, data := range datas {
		var category dto.FoodCategory
		if err := util.Parser(data, &category); err != nil {
			return nil, f.errHandler.Set(c, "parser error", err)
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
