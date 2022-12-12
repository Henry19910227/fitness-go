package diet

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateDiet(ctx *gin.Context)
	GetDiet(ctx *gin.Context)
}
