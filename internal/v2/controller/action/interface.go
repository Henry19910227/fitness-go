package action

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateCMSAction(ctx *gin.Context)
	UpdateCMSAction(ctx *gin.Context)
	GetCMSActions(ctx *gin.Context)

	CreateUserAction(ctx *gin.Context)
	UpdateUserAction(ctx *gin.Context)
	GetUserActions(ctx *gin.Context)
	DeleteUserAction(ctx *gin.Context)
	DeleteUserActionVideo(ctx *gin.Context)

	CreateTrainerAction(ctx *gin.Context)
	GetTrainerActions(ctx *gin.Context)
}
