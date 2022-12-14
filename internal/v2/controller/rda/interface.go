package rda

import "github.com/gin-gonic/gin"

type Controller interface {
	CalculateRDA(ctx *gin.Context)
	UpdateRDA(ctx *gin.Context)
}
