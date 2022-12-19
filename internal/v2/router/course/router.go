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

	v2.POST("/fcm_test", controller.FcmTest)

	v2.GET("/favorite/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetFavoriteCourses)

	v2.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	v2.GET("/cms/course/:course_id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
	v2.PATCH("/cms/courses/course_status", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesStatus)
	v2.PATCH("/cms/course/:course_id/cover", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesCover)

	v2.GET("/user/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourses)
	v2.POST("/user/course", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserCourse)
	v2.DELETE("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserCourse)
	v2.PATCH("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserCourse)
	v2.GET("/user/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourse)
	v2.GET("/user/course/:course_id/structure", midd.Verify([]global.Role{global.UserRole}), controller.GetUserCourseStructure)

	v2.GET("/trainer/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourses)
	v2.GET("/trainer/course/:course_id/overview", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourseOverview)
	v2.GET("/trainer/course/:course_id/statistic", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourseStatistic)
	v2.GET("/trainer/course/statistics", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourseStatistics)
	v2.POST("/trainer/course", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateTrainerCourse)
	v2.DELETE("/trainer/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteTrainerCourse)
	v2.GET("/trainer/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourse)
	v2.PATCH("/trainer/course/:course_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.UpdateTrainerCourse)
	v2.POST("/trainer/course/:course_id/submit", midd.Verify([]global.Role{global.UserRole}), controller.SubmitTrainerCourse)

	v2.GET("/store/course/:course_id", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreCourse)
	v2.GET("/store/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreCourses)
	v2.GET("/store/course/:course_id/structure", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreCourseStructure)
	v2.GET("/store/trainer/:user_id/courses", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreTrainerCourses)
	v2.GET("/store/home_page", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreHomePage)

}
