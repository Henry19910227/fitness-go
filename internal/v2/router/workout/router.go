package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := workout.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/workout/start_audio", http.Dir("./volumes/storage/workout/start_audio"))
	v2.StaticFS("/resource/workout/end_audio", http.Dir("./volumes/storage/workout/end_audio"))
	v2.POST("/user/plan/:plan_id/workout", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkout)
	v2.DELETE("/user/workout/:workout_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkout)
	v2.GET("/user/plan/:plan_id/workouts", midd.Verify([]global.Role{global.UserRole}), controller.GetUserWorkouts)
	v2.PATCH("/user/workout/:workout_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserWorkout)
	v2.DELETE("/user/workout/:workout_id/start_audio", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutStartAudio)
	v2.DELETE("/user/workout/:workout_id/end_audio", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutEndAudio)
}
