package user

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetCMSCourseUsers(ctx *gin.Context)

	UpdatePassword(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
	UpdateUserAvatar(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	GetAppleRefreshToken(ctx *gin.Context)
	RegisterForEmail(ctx *gin.Context)
	RegisterForFacebook(ctx *gin.Context)
	RegisterForGoogle(ctx *gin.Context)
	RegisterForLine(ctx *gin.Context)
	RegisterForApple(ctx *gin.Context)
	LoginForEmail(ctx *gin.Context)
	LoginForFacebook(ctx *gin.Context)
	LoginForGoogle(ctx *gin.Context)
	LoginForLine(ctx *gin.Context)
	LoginForApple(ctx *gin.Context)
	Logout(ctx *gin.Context)
	CreateRegisterOTP(ctx *gin.Context)
	RegisterEmailAccountValidate(ctx *gin.Context)
	RegisterFacebookAccountValidate(ctx *gin.Context)
	RegisterLineAccountValidate(ctx *gin.Context)
	RegisterGoogleAccountValidate(ctx *gin.Context)
	RegisterAppleAccountValidate(ctx *gin.Context)
	RegisterNicknameValidate(ctx *gin.Context)
	RegisterEmailValidate(ctx *gin.Context)
	CreateResetOTP(ctx *gin.Context)
	ResetOTPValidate(ctx *gin.Context)
	UpdateResetPassword(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}
