package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/action"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := action.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	v2.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	v2.GET("/cms/actions", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSActions)
	v2.POST("/cms/action", midd.Verify([]global.Role{global.AdminRole}), controller.CreateCMSAction)
}
