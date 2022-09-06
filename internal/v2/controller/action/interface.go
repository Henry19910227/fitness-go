package action

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserAction(ctx *gin.Context)
	UpdateUserAction(ctx *gin.Context)
	GetUserActions(ctx *gin.Context)
	DeleteUserActionVideo(ctx *gin.Context)
	GetCMSActions(ctx *gin.Context)
	CreateCMSAction(ctx *gin.Context)
	UpdateCMSAction(ctx *gin.Context)
}
