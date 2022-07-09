package banner

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateCMSBanner(ctx *gin.Context)
}
