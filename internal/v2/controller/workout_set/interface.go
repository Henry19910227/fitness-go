package workout_set

import "github.com/gin-gonic/gin"

type Controller interface {
	GetCMSWorkoutSets(ctx *gin.Context)
}
