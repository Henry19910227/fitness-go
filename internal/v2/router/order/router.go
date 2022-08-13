package order

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/order"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := order.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/course_order", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateCourseOrder)
	v2.POST("/subscribe_order", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateSubscribeOrder)
	v2.POST("/verify_apple_receipt", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.VerifyAppleReceipt)
	v2.POST("/verify_apple_subscribe", midd.Verify([]global.Role{global.UserRole}), controller.VerifyAppleSubscribe)
	v2.POST("/verify_google_subscribe", midd.Verify([]global.Role{global.UserRole}), controller.VerifyGoogleReceipt)
	v2.POST("/app_store_notification/v2", middleware.Transaction(orm.Shared().DB()), controller.AppStoreNotification)
	v2.GET("/cms/orders", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSOrders)
}
