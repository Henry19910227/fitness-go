package user

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdatePassword(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
	UpdateUserAvatar(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	RegisterForEmail(ctx *gin.Context)
	RegisterForFacebook(ctx *gin.Context)
	LoginForEmail(ctx *gin.Context)
	LoginForFacebook(ctx *gin.Context)
	Logout(ctx *gin.Context)
	CreateRegisterOTP(ctx *gin.Context)
	RegisterEmailAccountValidate(ctx *gin.Context)
	RegisterFacebookAccountValidate(ctx *gin.Context)
	RegisterNicknameValidate(ctx *gin.Context)
	RegisterEmailValidate(ctx *gin.Context)
}
