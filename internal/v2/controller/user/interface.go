package user

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdatePassword(ctx *gin.Context)
	RegisterForEmail(ctx *gin.Context)
	LoginForEmail(ctx *gin.Context)
	Logout(ctx *gin.Context)
	CreateRegisterOTP(ctx *gin.Context)
	RegisterAccountValidate(ctx *gin.Context)
	RegisterNicknameValidate(ctx *gin.Context)
}
