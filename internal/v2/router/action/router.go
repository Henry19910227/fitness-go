package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/action"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := action.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	v2.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	v2.POST("/user/action", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserAction)
	v2.GET("/cms/actions", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSActions)
	v2.POST("/cms/action", midd.Verify([]global.Role{global.AdminRole}), controller.CreateCMSAction)
	v2.PATCH("/cms/action/:action_id", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSAction)
}
