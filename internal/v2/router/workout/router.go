package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := workout.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/user/plan/:plan_id/workout", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkout)
	v2.DELETE("/user/workout/:workout_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkout)
}
