package admin

import "github.com/gin-gonic/gin"

type Controller interface {
	CMSLogin(ctx *gin.Context)
}
