package course

import (
	"github.com/Henry19910227/fitness-go/internal/controller/course"
	"github.com/Henry19910227/fitness-go/internal/global"
	middleware "github.com/Henry19910227/fitness-go/internal/middleware/token"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetRoute(route *gin.Engine, gormTool tool.Gorm, redisTool tool.Redis, viperTool *viper.Viper) *gin.Engine {
	controller := course.NewController(gormTool)
	midd := middleware.NewTokenMiddleware(redisTool, viperTool)
	v1 := route.Group("/api/v1")
	v1.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	v1.GET("/cms/course/:id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
	return route
}
