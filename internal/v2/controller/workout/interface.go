package workout

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserWorkout(ctx *gin.Context)
	DeleteUserWorkout(ctx *gin.Context)
	GetUserWorkouts(ctx *gin.Context)
}
