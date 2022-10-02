package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/controller/workout_set"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := workoutSet.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/workout_set/start_audio", http.Dir("./volumes/storage/workout_set/start_audio"))
	v2.StaticFS("/resource/workout_set/progress_audio", http.Dir("./volumes/storage/workout_set/progress_audio"))

	v2.GET("/cms/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSWorkoutSets)

	v2.POST("/user/workout/:workout_id/workout_sets", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkoutSets)
	v2.POST("/user/workout_set/:workout_set_id/duplicate", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserWorkoutSetByDuplicate)
	v2.POST("/user/workout/:workout_id/rest_set", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserRestSet)
	v2.DELETE("/user/workout_set/:workout_set_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutSet)
	v2.PATCH("/user/workout_set/:workout_set_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserWorkoutSet)
	v2.DELETE("/user/workout_set/:workout_set_id/start_audio", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutSetStartAudio)
	v2.DELETE("/user/workout_set/:workout_set_id/progress_audio", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserWorkoutSetProgressAudio)
	v2.GET("/user/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.UserRole}), controller.GetUserWorkoutSets)

	v2.GET("/trainer/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerWorkoutSets)
}
