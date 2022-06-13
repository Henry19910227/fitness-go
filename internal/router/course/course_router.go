package course

import (
	"github.com/Henry19910227/fitness-go/internal/controller/course"
	"github.com/Henry19910227/fitness-go/internal/global"
	middleware "github.com/Henry19910227/fitness-go/internal/middleware/token"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetRoute(baseGroup *gin.RouterGroup, gormTool tool.Gorm, redisTool tool.Redis, viperTool *viper.Viper) {
	controller := course.NewController(gormTool)
	midd := middleware.NewTokenMiddleware(redisTool, viperTool)
	baseGroup.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	baseGroup.GET("/cms/course/:id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
}
