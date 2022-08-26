package workout_set

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserWorkoutSets(ctx *gin.Context)
	GetCMSWorkoutSets(ctx *gin.Context)
}
