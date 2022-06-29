package body_image

import "github.com/gin-gonic/gin"

type Controller interface {
	GetBodyImages(ctx *gin.Context)
	CreateBodyImage(ctx *gin.Context)
}
