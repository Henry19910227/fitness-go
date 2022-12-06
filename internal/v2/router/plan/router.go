package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := plan.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.GET("/cms/course/:course_id/plans", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSPlans)

	v2.POST("/user/course/:course_id/plan", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateUserPlan)
	v2.DELETE("/user/plan/:plan_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteUserPlan)
	v2.GET("/user/course/:course_id/plans", midd.Verify([]global.Role{global.UserRole}), controller.GetUserPlans)
	v2.PATCH("/user/plan/:plan_id", midd.Verify([]global.Role{global.UserRole}), controller.UpdateUserPlan)

	v2.POST("/trainer/course/:course_id/plan", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateTrainerPlan)
	v2.GET("/trainer/course/:course_id/plans", midd.Verify([]global.Role{global.UserRole}), controller.GetTrainerPlans)
	v2.DELETE("/trainer/plan/:plan_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteTrainerPlan)
	v2.PATCH("/trainer/plan/:plan_id", midd.Verify([]global.Role{global.UserRole}), controller.UpdateTrainerPlan)

	v2.GET("/store/course/:course_id/plans", midd.Verify([]global.Role{global.UserRole}), controller.GetStorePlans)
}
