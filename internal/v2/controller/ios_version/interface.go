package ios_version

import "github.com/gin-gonic/gin"

type Controller interface {
	GetIOSVersion(ctx *gin.Context)
}
