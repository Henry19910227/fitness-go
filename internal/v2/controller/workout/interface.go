package workout

import "github.com/gin-gonic/gin"

type Controller interface {
	CreatePersonalWorkout(ctx *gin.Context)
	DeletePersonalWorkout(ctx *gin.Context)
}
