package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/controller/workout_set"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetRoute(baseGroup *gin.RouterGroup, gormTool tool.Gorm, redisTool tool.Redis, viperTool *viper.Viper) {
	controller := workoutSet.NewController(gormTool)
	midd := middleware.NewTokenMiddleware(redisTool, viperTool)
	baseGroup.GET("/cms/workout/:workout_id/workout_sets", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSWorkoutSets)
}
