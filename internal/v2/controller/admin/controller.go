package admin

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/admin/api_cms_login"
	"github.com/Henry19910227/fitness-go/internal/v2/model/admin/api_cms_logout"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver admin.Resolver
}

func New(resolver admin.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CMSLogin 管理者登入
// @Summary 管理者登入
// @Description 管理者登入
// @Tags CMS登入_v2
// @Accept json
// @Produce json
// @Param json_body body api_cms_login.Body true "輸入參數"
// @Success 200 {object} api_cms_login.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/login [POST]
func (c *controller) CMSLogin(ctx *gin.Context) {
	input := api_cms_login.Input{}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICMSLogin(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// CMSLogout 管理者登出
// @Summary 管理者登出
// @Description 管理者登出
// @Tags CMS登入_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} api_cms_logout.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/logout [POST]
func (c *controller) CMSLogout(ctx *gin.Context) {
	input := api_cms_logout.Input{}
	input.ID = ctx.MustGet("uid").(int64)
	output := c.resolver.APICMSLogout(&input)
	ctx.JSON(http.StatusOK, output)
}
