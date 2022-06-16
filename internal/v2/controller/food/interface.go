package food

import "github.com/gin-gonic/gin"

type Controller interface {
	GetFoods(ctx *gin.Context)
}
