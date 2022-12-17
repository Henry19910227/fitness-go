package android_version

import "github.com/gin-gonic/gin"

type Controller interface {
	GetAndroidVersion(ctx *gin.Context)
}
