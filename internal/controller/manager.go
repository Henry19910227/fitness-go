package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Manager struct {
	Base
	loginService service.Login
}

func NewManagerController(
	baseGroup *gin.RouterGroup,
	loginService service.Login,
	middleware gin.HandlerFunc) {

	manager := &Manager{
		loginService: loginService,
	}

	//不需驗證token
	baseGroup.POST("/manager/login", manager.Login)

	//需驗證token
	managerGroup := baseGroup.Group("/manager")
	managerGroup.Use(middleware)
	managerGroup.POST("/logout", manager.Logout)
}

// Login 管理員登入
// @Summary 管理員登入
// @Description 管理員登入
// @Tags Manager
// @Accept json
// @Produce json
// @Param json_body body validator.AdminLoginBody true "輸入參數"
// @Success 200 {object} model.SuccessLoginResult{data=admindata.Admin} "登入成功"
// @Failure 400 {object} model.ErrorResult "登入失敗"
// @Router /manager/login [POST]
func (m *Manager) Login(c *gin.Context) {
	var body validator.AdminLoginBody
	if err := c.ShouldBindJSON(&body); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	admin, token, err := m.loginService.LoginForAdmin(c, body.Email, body.Password)
	if err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONLoginSuccessResponse(c, token, admin, "admin login success!")
}

// Logout 管理員登出
// @Summary 管理員登出
// @Description 管理員登出
// @Tags Manager
// @Accept json
// @Produce json
// @Security icebaby_admin_token
// @Success 200 {object} model.SuccessResult "登出成功"
// @Failure 400 {object} model.ErrorResult "登出失敗"
// @Router /manager/logout [POST]
func (m *Manager) Logout(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		m.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := m.loginService.LogoutForAdmin(c, header.Token); err != nil {
		m.JSONErrorResponse(c, err)
		return
	}
	m.JSONSuccessResponse(c, nil, "登出成功")
}