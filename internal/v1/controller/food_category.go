package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/gin-gonic/gin"
)

type FoodCategory struct {
	Base
	foodCategoryService service.FoodCategory
}

func NewFoodCategory(baseGroup *gin.RouterGroup, foodCategoryService service.FoodCategory, userMiddleware middleware.User) {
	foodCategory := &FoodCategory{foodCategoryService: foodCategoryService}
	baseGroup.GET("/food_categories",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		foodCategory.GetFoodCategories)
}

// GetFoodCategories 獲取食物分類
// @Summary 獲取食物分類
// @Description 獲取食物分類
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=[]dto.FoodCategory} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/food_categories [GET]
func (f *FoodCategory) GetFoodCategories(c *gin.Context) {
	categories, err := f.foodCategoryService.GetFoodCategories(c)
	if err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, categories, "success!")
}
