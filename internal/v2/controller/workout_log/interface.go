package workout_log

import "github.com/gin-gonic/gin"

type Controller interface {
	CreateUserWorkoutLog(ctx *gin.Context)
	GetUserWorkoutLogs(ctx *gin.Context)
	GetUserWorkoutLog(ctx *gin.Context)
	DeleteUserWorkoutLog(ctx *gin.Context)
}
