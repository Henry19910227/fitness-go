package android_version

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/android_version/api_get_android_version"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/android_version"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver android_version.Resolver
}

func New(resolver android_version.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetAndroidVersion 獲取 android 版本號
// @Summary 獲取 android 版本號
// @Description 獲取 android 版本號
// @Tags 版本管理_v2
// @Accept json
// @Produce json
// @Success 200 {object} api_get_android_version.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/android_version [GET]
func (c *controller) GetAndroidVersion(ctx *gin.Context) {
	input := api_get_android_version.Input{}
	output := c.resolver.APIGetAndroidVersion(&input)
	ctx.JSON(http.StatusOK, output)
}
