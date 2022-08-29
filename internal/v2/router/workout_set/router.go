package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/controller/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := workoutSet.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/user/workout/:workout_id/workout_sets", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkoutSets)
	v2.DELETE("/user/workout_set/:workout_set_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutSet)
	v2.GET("/cms/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSWorkoutSets)
}
