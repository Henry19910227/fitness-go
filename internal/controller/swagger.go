package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/gin-gonic/gin"
)

type SwaggerController struct {
	swagService service.Swagger
}

func NewSwaggerController(router *gin.Engine, swagService service.Swagger) {
	api := router.Group("/api")
	api.GET("/swagger/*any", swagService.WrapHandler())
}

