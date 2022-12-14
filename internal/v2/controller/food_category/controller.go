package food_category

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/food_category/api_get_food_categories"
	foodCategory "github.com/Henry19910227/fitness-go/internal/v2/resolver/food_category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver foodCategory.Resolver
}

func New(resolver foodCategory.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetFoodCategories 獲取食物分類
// @Summary 獲取食物分類
// @Description 獲取食物分類
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_get_food_categories.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/food_categories [GET]
func (c *controller) GetFoodCategories(ctx *gin.Context) {
	input := api_get_food_categories.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetFoodCategories(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSFoodCategories 獲取食物分類
// @Summary 獲取食物分類
// @Description 獲取食物分類
// @Tags CMS內容管理_食品庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} food_category.APIGetCMSFoodCategoriesOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/food_categories [GET]
func (c *controller) GetCMSFoodCategories(ctx *gin.Context) {
	output := c.resolver.APIGetCMSFoodCategories()
	ctx.JSON(http.StatusOK, output)
}
