package favorite_course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/favorite_course"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := favorite_course.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.POST("/favorite/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.CreateFavoriteCourse)
	v2.DELETE("/favorite/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteFavoriteCourse)
}
