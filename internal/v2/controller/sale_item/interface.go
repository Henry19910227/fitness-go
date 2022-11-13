package sale_item

import "github.com/gin-gonic/gin"

type Controller interface {
	GetSaleItems(ctx *gin.Context)
}
