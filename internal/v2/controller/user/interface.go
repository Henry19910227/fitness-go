package user

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdatePassword(ctx *gin.Context)
}
