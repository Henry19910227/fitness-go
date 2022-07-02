package food_category

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSFoodCategories(ctx *gin.Context)
}
