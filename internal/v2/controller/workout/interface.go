package workout

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserWorkout(ctx *gin.Context)
	DeleteUserWorkout(ctx *gin.Context)
	GetUserWorkouts(ctx *gin.Context)
	UpdateUserWorkout(ctx *gin.Context)
	DeleteUserWorkoutStartAudio(ctx *gin.Context)
}
