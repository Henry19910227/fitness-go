package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/gin-gonic/gin"
)

type meal struct {
	Base
	mealRepo   repository.Meal
	errHandler errcode.Handler
}

func NewMeal(mealRepo repository.Meal, errHandler errcode.Handler) Meal {
	return &meal{mealRepo: mealRepo, errHandler: errHandler}
}

func (m *meal) CreateMeals(c *gin.Context, param *dto.CreateMealsParam) errcode.Error {
	mealItems := make([]*model.MealItem, 0)
	if err := util.Parser(param.MealParamItems, &mealItems); err != nil {
		return m.errHandler.Set(c, "parser error", err)
	}
	_, err := m.mealRepo.SaveMeals(&model.SaveMealsParam{
		MealItems: mealItems,
	})
	if err != nil {
		return m.errHandler.Set(c, "meal repo", err)
	}
	return nil
}
