package diet

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet/api_create_diet"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/diet"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver diet.Resolver
}

func New(resolver diet.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateDiet 創建飲食紀錄
// @Summary 創建飲食紀錄
// @Description 創建飲食紀錄
// @Tags 飲食_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body api_create_diet.Body true "輸入參數"
// @Success 200 {object} api_create_diet.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/diet [POST]
func (c *controller) CreateDiet(ctx *gin.Context) {
	input := api_create_diet.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateDiet(&input)
	ctx.JSON(http.StatusOK, output)
}
