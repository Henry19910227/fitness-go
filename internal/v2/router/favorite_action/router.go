package favorite_action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/favorite_action"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := favorite_action.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/favorite/action/:action_id", midd.Verify([]global.Role{global.UserRole}), controller.CreateFavoriteAction)
	v2.DELETE("/favorite/action/:action_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteFavoriteAction)
}
