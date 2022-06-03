package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Food struct {
	Base
	foodService service.Food
}

func NewFood(baseGroup *gin.RouterGroup, foodService service.Food, userMiddleware middleware.User)  {
	food := &Food{foodService: foodService}
	baseGroup.POST("/food",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		food.CreateFood)
}

// CreateFood 創建食物
// @Summary 創建食物
// @Description 創建食物
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateFoodBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Food} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /food [POST]
func (f *Food) CreateFood(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.CreateFoodBody
	if err := c.ShouldBindJSON(&body); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	food, err := f.foodService.CreateFood(c, &dto.CreateFoodParam{
		UserID: uid,
		FoodCategoryID: body.FoodCategoryID,
		Source: 2,
		Name: body.Name,
		AmountDesc: body.AmountDesc,
	})
	if err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, food, "success!")
}
