package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/ios_version/api_get_ios_version"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/ios_version"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver ios_version.Resolver
}

func New(resolver ios_version.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetIOSVersion 獲取 ios 版本號
// @Summary 獲取 ios 版本號
// @Description 獲取 ios 版本號
// @Tags 版本管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_get_ios_version.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/ios_version [GET]
func (c *controller) GetIOSVersion(ctx *gin.Context) {
	input := api_get_ios_version.Input{}
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetIOSVersion(&input)
	ctx.JSON(http.StatusOK, output)
}