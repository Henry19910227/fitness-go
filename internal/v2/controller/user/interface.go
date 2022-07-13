package user

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdatePassword(ctx *gin.Context)
	RegisterAccountValidate(ctx *gin.Context)
	RegisterNicknameValidate(ctx *gin.Context)
}
