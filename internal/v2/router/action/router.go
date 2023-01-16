package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/action"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := action.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/action/cover", http.Dir("./volumes/storage/action/cover"))
	v2.StaticFS("/resource/action/video", http.Dir("./volumes/storage/action/video"))
	v2.StaticFS("/resource/action/system_image", http.Dir("./volumes/storage/action/system_image"))

	v2.GET("/cms/actions", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSActions)
	v2.POST("/cms/action", midd.Verify([]global.Role{global.AdminRole}), controller.CreateCMSAction)
	v2.PATCH("/cms/action/:action_id", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSAction)

	v2.POST("/user/action", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserAction)
	v2.PATCH("/user/action/:action_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserAction)
	v2.GET("/user/actions", midd.Verify([]global.Role{global.UserRole}), controller.GetUserActions)
	v2.GET("/user/action/:action_id/best_personal_record", midd.Verify([]global.Role{global.UserRole}), controller.GetUserActionBestPR)
	v2.DELETE("/user/action/:action_id", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserAction)
	v2.DELETE("/user/action/:action_id/video", midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserActionVideo)
	v2.GET("/user/action/system_images", midd.Verify([]global.Role{global.UserRole}), controller.APIGetUserActionSystemImages)

	v2.POST("/trainer/course/:course_id/action", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateTrainerAction)
	v2.PATCH("/trainer/action/:action_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.UpdateTrainerAction)
	v2.GET("/trainer/course/:course_id/actions", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerCourseActions)
	v2.DELETE("/trainer/action/:action_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteTrainerAction)
	v2.DELETE("/trainer/action/:action_id/video", midd.Verify([]global.Role{global.UserRole}), controller.DeleteTrainerActionVideo)
}
