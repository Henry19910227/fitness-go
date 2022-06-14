package controller

import (
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/gin-gonic/gin"
)

type SwaggerController struct {
	swagService service.Swagger
}

func NewSwagger(router *gin.Engine, swagService service.Swagger) {
	api := router.Group("/api")
	api.GET("/swagger/*any", swagService.WrapHandler())
}

