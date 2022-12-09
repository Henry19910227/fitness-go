package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food/api_get_foods"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/food"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver food.Resolver
}

func New(resolver food.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetFoods 獲取食物列表
// @Summary 獲取食物列表
// @Description 獲取食物列表
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param name query string false "食物名稱"
// @Param tag query int false "食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)"
// @Success 200 {object} api_get_foods.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/foods [GET]
func (c *controller) GetFoods(ctx *gin.Context) {
	input := api_get_foods.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetFoods(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetCMSFoods 獲取食物列表
// @Summary 獲取食物列表
// @Description 獲取食物列表
// @Tags CMS內容管理_食品庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} food.APIGetCMSFoodsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/foods [GET]
func (c *controller) GetCMSFoods(ctx *gin.Context) {
	input := model.APIGetCMSFoodsInput{}
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSFoods(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateCMSFood 創建食物
// @Summary 創建食物
// @Description 創建食物
// @Tags CMS內容管理_食品庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body food.APICreateCMSFoodBody true "輸入參數"
// @Success 200 {object} food.APICreateCMSFoodOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/food [POST]
func (c *controller) CreateCMSFood(ctx *gin.Context) {
	input := model.APICreateCMSFoodInput{}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateCMSFood(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSFood 修改食物
// @Summary 修改食物
// @Description 修改食物
// @Tags CMS內容管理_食品庫_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param food_id path int64 true "食物id"
// @Param json_body body food.APIUpdateCMSFoodBody true "輸入參數"
// @Success 200 {object} base.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/food/{food_id} [PATCH]
func (c *controller) UpdateCMSFood(ctx *gin.Context) {
	input := model.APIUpdateCMSFoodInput{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateCMSFood(&input)
	ctx.JSON(http.StatusOK, output)
}
