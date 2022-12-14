package food_category

import "github.com/gin-gonic/gin"

type Controller interface {
	GetFoodCategories(ctx *gin.Context)
	GetCMSFoodCategories(ctx *gin.Context)
}
