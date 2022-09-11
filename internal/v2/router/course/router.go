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

func SetRoute(v2 *gin.RouterGroup) {
	controller := course.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/course/cover", http.Dir("./volumes/storage/course/cover"))
	v2.GET("/favorite/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetFavoriteCourses)
	v2.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	v2.GET("/cms/course/:course_id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
	v2.PATCH("/cms/courses/course_status", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesStatus)
	v2.PATCH("/cms/course/:course_id/cover", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesCover)
	v2.GET("/user/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourses)
	v2.POST("/user/course", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserCourse)
	v2.DELETE("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserCourse)
	v2.PATCH("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserCourse)
	v2.GET("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourse)

	v2.GET("/trainer/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourses)
}
