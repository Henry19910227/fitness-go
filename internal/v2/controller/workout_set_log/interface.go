package workout_set_log

import "github.com/gin-gonic/gin"

type Controller interface {
	GetUserActionWorkoutSetLogs(ctx *gin.Context)
}
