package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
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

func (m *meal) DeleteMeal(c *gin.Context, mealID int64, userID int64) errcode.Error {
	ownerID, err := m.mealRepo.FindMealOwner(mealID)
	if err != nil {
		return m.errHandler.Set(c, "meal repo", err)
	}
	if ownerID != userID {
		return m.errHandler.PermissionDenied()
	}
	if err := m.mealRepo.DeleteMeal(mealID); err != nil {
		return m.errHandler.Set(c, "meal repo", err)
	}
	return nil
}
