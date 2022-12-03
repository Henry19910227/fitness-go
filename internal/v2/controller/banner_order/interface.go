package banner_order

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdateCMSBannerOrders(ctx *gin.Context)
}
