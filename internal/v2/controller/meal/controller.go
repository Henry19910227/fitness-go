package meal

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	mealModel "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
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

// UpdateMeals 修改餐食
// @Summary 修改餐食
// @Description 修改餐食
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param diet_id path int64 true "飲食id"
// @Param json_body body []meal.APIPutMealsInputItem true "輸入參數"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/diet/{diet_id}/meals [PUT]
func (c *controller) UpdateMeals(ctx *gin.Context) {
	var uri struct {
		dietModel.IDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	meals := make([]*mealModel.APIPutMealsInputItem, 0)
	if err := ctx.ShouldBindJSON(&meals); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	tx := ctx.MustGet("db_trx").(*gorm.DB)
	input := mealModel.APIPutMealsInput{}
	input.DietID = uri.ID
	input.Meals = meals
	if err := c.resolver.APIPutMeals(tx, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	ctx.JSON(http.StatusOK, baseModel.Success())
}
