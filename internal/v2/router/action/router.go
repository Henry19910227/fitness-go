package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/action"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(baseGroup *gin.RouterGroup) {
	controller := action.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	baseGroup.GET("/cms/actions", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSActions)
}
