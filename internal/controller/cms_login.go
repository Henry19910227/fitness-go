package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type CMSLogin struct {
	Base
	loginService service.Login
}

func NewCMSLogin(baseGroup *gin.RouterGroup, loginService service.Login, userMiddleware middleware.User) {
	cms := &CMSLogin{loginService: loginService}

	baseGroup.POST("/cms/login", cms.Login)
	baseGroup.POST("/cms/logout",
		userMiddleware.TokenPermission([]global.Role{global.AdminRole}),
		cms.Logout)
}

// Login 管理者登入
// @Summary 管理者登入
// @Description 管理者登入
// @Tags CMS/Login_v1
// @Accept json
// @Produce json
// @Param json_body body validator.CMSLoginByEmailBody true "輸入參數"
// @Success 200 {object} model.SuccessLoginResult{data=dto.Admin} "登入成功"
// @Failure 400 {object} model.ErrorResult "登入失敗"
// @Router /v1/cms/login [POST]
func (l *CMSLogin) Login(c *gin.Context) {
	var body validator.CMSLoginByEmailBody
	if err := c.ShouldBindJSON(&body); err != nil {
		l.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	user, token, err := l.loginService.AdminLoginByEmail(c, body.Email, body.Password)
	if err != nil {
		l.JSONErrorResponse(c, err)
		return
	}
	l.JSONLoginSuccessResponse(c, token, user, "login success!")
}

// Logout 管理者登出
// @Summary 管理者登出
// @Description 管理者登出
// @Tags CMS/Login_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult "登出成功"
// @Failure 400 {object} model.ErrorResult "登出失敗"
// @Router /v1/cms/logout [POST]
func (l *CMSLogin) Logout(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		l.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := l.loginService.AdminLogoutByToken(c, header.Token); err != nil {
		l.JSONErrorResponse(c, err)
		return
	}
	l.JSONSuccessResponse(c, nil, "登出成功")
}
