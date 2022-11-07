package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/workout_log"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := workout_log.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/user/workout/:workout_id/workout_log", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkoutLog)
	v2.GET("/user/workout_logs", midd.Verify([]global.Role{global.UserRole}), controller.GetUserWorkoutLogs)
	v2.GET("/user/workout_log/:workout_log_id", midd.Verify([]global.Role{global.UserRole}), controller.GetUserWorkoutLog)
	v2.DELETE("/user/workout_log/:workout_log_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutLog)
}
