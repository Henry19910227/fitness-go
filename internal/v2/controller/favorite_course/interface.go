package favorite_course

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateFavoriteCourse(ctx *gin.Context)
	DeleteFavoriteCourse(ctx *gin.Context)
}
