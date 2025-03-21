package workout

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserWorkout(ctx *gin.Context)
	DeleteUserWorkout(ctx *gin.Context)
	GetUserWorkouts(ctx *gin.Context)
	UpdateUserWorkout(ctx *gin.Context)
	DeleteUserWorkoutStartAudio(ctx *gin.Context)
	DeleteUserWorkoutEndAudio(ctx *gin.Context)

	CreateTrainerWorkout(ctx *gin.Context)
	DeleteTrainerWorkout(ctx *gin.Context)
	GetTrainerWorkouts(ctx *gin.Context)
	UpdateTrainerWorkout(ctx *gin.Context)
	DeleteTrainerWorkoutStartAudio(ctx *gin.Context)
	DeleteTrainerWorkoutEndAudio(ctx *gin.Context)

	GetStoreWorkouts(ctx *gin.Context)
}
