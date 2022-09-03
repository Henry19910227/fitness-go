package action

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserAction(ctx *gin.Context)
	GetCMSActions(ctx *gin.Context)
	CreateCMSAction(ctx *gin.Context)
	UpdateCMSAction(ctx *gin.Context)
}
