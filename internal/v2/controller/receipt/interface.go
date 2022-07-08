package receipt

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSOrderReceipts(ctx *gin.Context)
}
