package action

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateCMSAction(ctx *gin.Context)
	UpdateCMSAction(ctx *gin.Context)
	GetCMSActions(ctx *gin.Context)

	CreateUserAction(ctx *gin.Context)
	UpdateUserAction(ctx *gin.Context)
	GetUserActions(ctx *gin.Context)
	GetUserActionBestPR(ctx *gin.Context)
	DeleteUserAction(ctx *gin.Context)
	DeleteUserActionVideo(ctx *gin.Context)
	APIGetUserActionSystemImages(ctx *gin.Context)

	CreateTrainerAction(ctx *gin.Context)
	UpdateTrainerAction(ctx *gin.Context)
	GetTrainerCourseActions(ctx *gin.Context)
	DeleteTrainerAction(ctx *gin.Context)
	DeleteTrainerActionVideo(ctx *gin.Context)
}
