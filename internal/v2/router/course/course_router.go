package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/course"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(baseGroup *gin.RouterGroup) {
	controller := course.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	baseGroup.StaticFS("/resource/course/cover", http.Dir("./volumes/storage/course/cover"))
	baseGroup.GET("/favorite/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetFavoriteCourses)
	baseGroup.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	baseGroup.GET("/cms/course/:course_id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
	baseGroup.PATCH("/cms/courses/course_status", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesStatus)
	baseGroup.PATCH("/cms/course/:course_id/cover", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesCover)
	baseGroup.GET("/user/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourses)
	baseGroup.POST("/user/course", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserCourse)
}
