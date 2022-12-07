package rda

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdateRDA(ctx *gin.Context)
}
