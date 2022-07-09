package banner

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateCMSBanner(ctx *gin.Context)
	GetCMSBanners(ctx *gin.Context)
	DeleteCMSBanner(ctx *gin.Context)
}
