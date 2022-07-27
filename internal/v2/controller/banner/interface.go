package banner

import "github.com/gin-gonic/gin"

type Controller interface {
	GetBanners(ctx *gin.Context)
	CreateCMSBanner(ctx *gin.Context)
	GetCMSBanners(ctx *gin.Context)
	DeleteCMSBanner(ctx *gin.Context)
}
