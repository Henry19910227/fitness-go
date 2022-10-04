package workout_set

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSWorkoutSets(ctx *gin.Context)

	CreateUserWorkoutSets(ctx *gin.Context)
	CreateUserWorkoutSetByDuplicate(ctx *gin.Context)
	CreateUserRestSet(ctx *gin.Context)
	DeleteUserWorkoutSet(ctx *gin.Context)
	UpdateUserWorkoutSet(ctx *gin.Context)
	DeleteUserWorkoutSetStartAudio(ctx *gin.Context)
	DeleteUserWorkoutSetProgressAudio(ctx *gin.Context)
	GetUserWorkoutSets(ctx *gin.Context)

	CreateTrainerWorkoutSets(ctx *gin.Context)
	CreateTrainerWorkoutSetByDuplicate(ctx *gin.Context)
	GetTrainerWorkoutSets(ctx *gin.Context)
	DeleteTrainerWorkoutSet(ctx *gin.Context)
	UpdateTrainerWorkoutSet(ctx *gin.Context)
	DeleteTrainerWorkoutSetStartAudio(ctx *gin.Context)
	DeleteTrainerWorkoutSetProgressAudio(ctx *gin.Context)
	CreateTrainerRestSet(ctx *gin.Context)
}
