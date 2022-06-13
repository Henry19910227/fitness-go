package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Base
	loginService service.Login
}

func NewLogin(baseGroup *gin.RouterGroup, loginService service.Login, userMiddle gin.HandlerFunc, adminMiddle gin.HandlerFunc) {
	login := &Login{
		loginService: loginService,
	}
	baseGroup.POST("/login/user/email", login.UserLoginByEmail)
	baseGroup.POST("/login/admin/email", login.AdminLoginByEmail)

	//驗證 user token
	UserGroup := baseGroup.Group("/logout/user")
	UserGroup.Use(userMiddle)
	UserGroup.POST("", login.UserLogout)

	//驗證 admin token
	adminLogoutGroup := baseGroup.Group("/logout/admin")
	adminLogoutGroup.Use(adminMiddle)
	adminLogoutGroup.POST("", login.AdminLogout)
}

// UserLoginByEmail 用戶使用信箱登入
// @Summary 用戶使用信箱登入
// @Description 用戶使用信箱登入
// @Tags Login_v1
// @Accept json
// @Produce json
// @Param json_body body validator.LoginByEmailBody true "輸入參數"
// @Success 200 {object} model.SuccessLoginResult{data=dto.User} "登入成功"
// @Failure 400 {object} model.ErrorResult "登入失敗"
// @Router /v1/login/user/email [POST]
func (l *Login) UserLoginByEmail(c *gin.Context) {
	var body validator.LoginByEmailBody
	if err := c.ShouldBindJSON(&body); err != nil {
		l.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	user, token, err := l.loginService.UserLoginByEmail(c, body.Email, body.Password)
	if err != nil {
		l.JSONErrorResponse(c, err)
		return
	}
	l.JSONLoginSuccessResponse(c, token, user, "login success!")
}

// AdminLoginByEmail 管理者使用信箱登入
// @Summary 管理者使用信箱登入
// @Description 管理者使用信箱登入
// @Tags Login_v1
// @Accept json
// @Produce json
// @Param json_body body validator.LoginByEmailBody true "輸入參數"
// @Success 200 {object} model.SuccessLoginResult{data=dto.Admin} "登入成功"
// @Failure 400 {object} model.ErrorResult "登入失敗"
// @Router /v1/login/admin/email [POST]
func (l *Login) AdminLoginByEmail(c *gin.Context) {
	var body validator.LoginByEmailBody
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

// AdminLogout 管理員登出
// @Summary 管理員登出
// @Description 管理員登出
// @Tags Login_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult "登出成功"
// @Failure 400 {object} model.ErrorResult "登出失敗"
// @Router /v1/logout/admin [POST]
func (l *Login) AdminLogout(c *gin.Context) {
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

// UserLogout 用戶登出
// @Summary 用戶登出
// @Description 用戶登出
// @Tags Login_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult "登出成功"
// @Failure 400 {object} model.ErrorResult "登出失敗"
// @Router /v1/logout/user [POST]
func (l *Login) UserLogout(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		l.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := l.loginService.UserLogoutByToken(c, header.Token); err != nil {
		l.JSONErrorResponse(c, err)
		return
	}
	l.JSONSuccessResponse(c, nil, "登出成功")
}
