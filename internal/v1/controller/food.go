package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type Food struct {
	Base
	foodService service.Food
}

func NewFood(baseGroup *gin.RouterGroup, foodService service.Food, userMiddleware middleware.User) {
	food := &Food{foodService: foodService}
	baseGroup.POST("/food",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		food.CreateFood)
	baseGroup.GET("/foods",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		food.GetFoods)
	baseGroup.GET("/recent_foods",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		food.GetRecentFoods)
	baseGroup.DELETE("/food/:food_id",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		food.DeleteFood)
}

// CreateFood 創建食物
// @Summary 創建食物
// @Description 創建食物
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateFoodBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Food} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/food [POST]
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
		UserID:         uid,
		FoodCategoryID: body.FoodCategoryID,
		Source:         2,
		Name:           body.Name,
		AmountDesc:     body.AmountDesc,
	})
	if err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, food, "success!")
}

// GetFoods 獲取食物列表
// @Summary 獲取食物列表 (API已經過時，更新為  /v2/foods [GET])
// @Description 獲取食物列表
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param food_category_tag query int true "食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)"
// @Success 200 {object} model.SuccessResult{data=[]dto.Food} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/foods [GET]
func (f *Food) GetFoods(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var query validator.GetFoodsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	foods, err := f.foodService.GetFoods(c, uid, global.FoodCategoryTag(query.FoodCategoryTag))
	if err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, foods, "success!")
}

// GetRecentFoods 獲取食物歷程列表
// @Summary 獲取食物歷程列表 (API已經過時，更新為 /v2/meals [GET])
// @Description 獲取食物歷程列表
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=[]dto.RecentFood} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/recent_foods [GET]
func (f *Food) GetRecentFoods(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	foods, err := f.foodService.GetRecentFood(c, uid)
	if err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, foods, "success!")
}

// DeleteFood 刪除食物
// @Summary 刪除食物
// @Description 刪除食物
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param food_id path int64 true "食物id"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/food/{food_id} [DELETE]
func (f *Food) DeleteFood(c *gin.Context) {
	uid, e := f.GetUID(c)
	if e != nil {
		f.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var uri validator.FoodIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		f.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := f.foodService.DeleteFood(c, uri.FoodID, uid); err != nil {
		f.JSONErrorResponse(c, err)
		return
	}
	f.JSONSuccessResponse(c, nil, "success!")
}
