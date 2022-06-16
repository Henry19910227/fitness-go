package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/food"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetRoute(baseGroup *gin.RouterGroup, gormTool orm.Tool, redisTool tool.Redis, viperTool *viper.Viper) {
	controller := food.NewController(gormTool, viperTool)
	midd := middleware.NewTokenMiddleware(redisTool, viperTool)
	baseGroup.GET("/foods", midd.Verify([]global.Role{global.UserRole}), controller.GetFoods)
}
