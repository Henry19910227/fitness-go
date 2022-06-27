package body_image

import "github.com/gin-gonic/gin"

type Controller interface {
	APIGetBodyImages(ctx *gin.Context)
}
