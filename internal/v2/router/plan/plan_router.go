package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/plan"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetRoute(baseGroup *gin.RouterGroup, gormTool tool.Gorm, redisTool tool.Redis, viperTool *viper.Viper) {
	controller := plan.NewController(gormTool)
	midd := middleware.NewTokenMiddleware(redisTool, viperTool)
	baseGroup.GET("/cms/course/:course_id/plans", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSPlans)
}
