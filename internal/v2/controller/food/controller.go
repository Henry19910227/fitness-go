package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	foodModel "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	foodCategoryModel "github.com/Henry19910227/fitness-go/internal/v2/model/food_category"
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
// @Success 200 {object} food.APIGetFoodsOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/foods [GET]
func (c *controller) GetFoods(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	var query struct {
		foodModel.NameField
		foodCategoryModel.TagField
	}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := foodModel.APIGetFoodsInput{}
	input.UserID = util.PointerInt64(uid.(int64))
	if err := util.Parser(query, &input); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetFoods(ctx, &input)
	ctx.JSON(http.StatusOK, output)
}
