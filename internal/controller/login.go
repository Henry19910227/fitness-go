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

func NewLogin(baseGroup *gin.RouterGroup, loginService service.Login)  {
	login := &Login{
		loginService: loginService,
	}
	baseGroup.POST("/login/user/email", login.UserLoginByEmail)
}

// UserLoginByEmail 用戶使用信箱登入
// @Summary 用戶使用信箱登入
// @Description 用戶使用信箱登入
// @Tags Login
// @Accept json
// @Produce json
// @Param json_body body validator.UserLoginByEmailBody true "輸入參數"
// @Success 200 {object} model.SuccessLoginResult{data=logindto.User} "登入成功"
// @Failure 400 {object} model.ErrorResult "登入失敗"
// @Router /login/user/email [POST]
func (l *Login) UserLoginByEmail(c *gin.Context)  {
	var body validator.UserLoginByEmailBody
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