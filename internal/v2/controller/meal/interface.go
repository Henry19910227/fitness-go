package meal

import "github.com/gin-gonic/gin"

type Controller interface {
	UpdateMeals(ctx *gin.Context)
	GetMeals(ctx *gin.Context)
}
