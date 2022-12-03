package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/banner_order"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := banner_order.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.PUT("/cms/banner_orders", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSBannerOrders)
}

