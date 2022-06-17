package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/food"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(baseGroup *gin.RouterGroup) {
	controller := food.NewController(orm.Shared())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	baseGroup.GET("/foods", midd.Verify([]global.Role{global.UserRole}), controller.GetFoods)
}
