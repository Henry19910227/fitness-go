package ios_version

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/ios_version"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := ios_version.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.GET("/ios_version", midd.Verify([]global.Role{global.UserRole}), controller.GetIOSVersion)
}
