package order

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSOrders(ctx *gin.Context)
}
