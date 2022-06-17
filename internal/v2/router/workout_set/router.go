package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/controller/workout_set"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(baseGroup *gin.RouterGroup) {
	controller := workoutSet.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	baseGroup.GET("/cms/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSWorkoutSets)
}
