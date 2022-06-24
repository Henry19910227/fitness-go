package action

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSActions(ctx *gin.Context)
	CreateCMSAction(ctx *gin.Context)
}
