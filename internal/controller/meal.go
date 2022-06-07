package controller

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/util"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Meal struct {
	Base
	mealService service.Meal
}

func NewMeal(baseGroup *gin.RouterGroup, mealService service.Meal, userMiddleware middleware.User)  {
	meal := &Meal{mealService: mealService}
	baseGroup.POST("/meals",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		meal.CreateMeals)
	baseGroup.DELETE("/meal/:meal_id",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		meal.DeleteMeal)
}

// CreateMeals 創建餐食
// @Summary 創建餐食
// @Description 創建餐食
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body []validator.MealParamItem true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /meals [POST]
func (m *Meal) CreateMeals(c *gin.Context) {
	datas := make([]*validator.MealParamItem, 0)
	if err := c.ShouldBindJSON(&datas); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	mealParamItems := make([]*dto.MealParamItem, 0)
	if err := util.Parser(datas, &mealParamItems); err != nil {
		m.JSONErrorResponse(c, errcode.NewError(8999, err))
		return
	}
	if err := m.mealService.CreateMeals(c, &dto.CreateMealsParam{
		MealParamItems: mealParamItems,
	}); err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, nil, "success!")
}

// DeleteMeal 刪除餐食
// @Summary 刪除餐食
// @Description 刪除餐食
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param meal_id path int64 true "餐食id"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /meal/{meal_id} [DELETE]
func (m *Meal) DeleteMeal(c *gin.Context) {
	uid, e := m.GetUID(c)
	if e != nil {
		m.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.MealIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := m.mealService.DeleteMeal(c, uri.MealID, uid); err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, nil, "success!")
}
