package user_subscribe_info

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_subscribe_info"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user_subscribe_info.Resolver
}

func New(resolver user_subscribe_info.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetUserSubscribeInfo 獲取用戶訂閱資訊
// @Summary 獲取用戶訂閱資訊
// @Description 獲取用戶訂閱資訊
// @Tags 用戶個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} user_subscribe_info.APIGetUserSubscribeInfoOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/subscribe_info [GET]
func (c *controller) GetUserSubscribeInfo(ctx *gin.Context) {
	var input model.APIGetUserSubscribeInfoInput
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetUserSubscribeInfo(&input)
	ctx.JSON(http.StatusOK, output)
}