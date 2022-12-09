package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
	"github.com/Henry19910227/fitness-go/internal/v2/model/meal/api_put_meals"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/meal"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver meal.Resolver
}

func New(resolver meal.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdateMeals 修改並覆蓋餐食
// @Summary 修改並覆蓋餐食
// @Description 修改並覆蓋餐食
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param diet_id path int64 true "飲食id"
// @Param json_body body api_put_meals.Body true "輸入參數"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/diet/{diet_id}/meals [PUT]
func (c *controller) UpdateMeals(ctx *gin.Context) {
	input := api_put_meals.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusOK, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusOK, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIPutMeals(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetMeals 獲取餐食列表
// @Summary 獲取餐食列表
// @Description 獲取餐食列表
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} meal.APIGetMealsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/meals [GET]
func (c *controller) GetMeals(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := mealModel.APIGetMealsInput{}
	input.UserID = util.PointerInt64(uid.(int64))
	output := c.resolver.APIGetMeals(&input)
	ctx.JSON(http.StatusOK, output)
}
