package user_course_asset

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_create_cms_course_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset/api_delete_cms_course_user"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user_course_asset"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user_course_asset.Resolver
}

func New(resolver user_course_asset.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateCMSCourseUsers 創建課表使用者
// @Summary 創建課表使用者
// @Description 創建課表使用者
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表 id"
// @Param json_body body api_create_cms_course_users.Body true "輸入參數"
// @Success 200 {object} api_create_cms_course_users.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id}/users [POST]
func (c *controller) CreateCMSCourseUsers(ctx *gin.Context) {
	input := api_create_cms_course_users.Input{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateCMSCourseUsers(&input)
	ctx.JSON(http.StatusOK, output)
}

// DeleteCMSCourseUser 刪除課表使用者
// @Summary 刪除課表使用者
// @Description 刪除課表使用者
// @Tags CMS課表管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id path int64 true "課表 id"
// @Param user_id path int64 true "用戶 id"
// @Success 200 {object} api_delete_cms_course_user.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/course/{course_id}/user/{user_id} [DELETE]
func (c *controller) DeleteCMSCourseUser(ctx *gin.Context) {
	input := api_delete_cms_course_user.Input{}
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIDeleteCMSCourseUser(&input)
	ctx.JSON(http.StatusOK, output)
}
