package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/food"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := food.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.GET("/foods", midd.Verify([]global.Role{global.UserRole}), controller.GetFoods)
	v2.GET("/cms/foods", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSFoods)
	v2.POST("/cms/food", midd.Verify([]global.Role{global.AdminRole}), controller.CreateCMSFood)
	v2.PATCH("/cms/food/:food_id", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSFood)
}
