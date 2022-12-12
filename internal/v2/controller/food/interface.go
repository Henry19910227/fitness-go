package food

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateFood(ctx *gin.Context)
	GetFoods(ctx *gin.Context)
	DeleteFood(ctx *gin.Context)

	GetCMSFoods(ctx *gin.Context)
	CreateCMSFood(ctx *gin.Context)
	UpdateCMSFood(ctx *gin.Context)
}
