package receipt

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSReceipts(ctx *gin.Context)
}
