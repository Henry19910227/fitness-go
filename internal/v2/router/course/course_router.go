package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/course"
	middleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(baseGroup *gin.RouterGroup) {
	controller := course.NewController(orm.Shared().DB())
	midd := middleware.NewTokenMiddleware(redis.Shared())
	baseGroup.StaticFS("/resource/course/cover", http.Dir("./volumes/storage/course/cover"))
	baseGroup.GET("/cms/courses", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourses)
	baseGroup.GET("/cms/course/:course_id", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourse)
	baseGroup.PATCH("/cms/courses/course_status", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesStatus)
	baseGroup.PATCH("/cms/course/:course_id/cover", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSCoursesCover)
}
