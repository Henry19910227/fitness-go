package order

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateCourseOrder(ctx *gin.Context)
	CreateSubscribeOrder(ctx *gin.Context)
	GetCMSOrders(ctx *gin.Context)
}
