package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/review"
	"github.com/Henry19910227/fitness-go/internal/v2/middleware"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := review.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.StaticFS("/resource/course/review", http.Dir("./volumes/storage/course/review"))
	v2.GET("/cms/reviews", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSReviews)
	v2.PATCH("/cms/review/:review_id", midd.Verify([]global.Role{global.AdminRole}), controller.UpdateCMSReview)
	v2.DELETE("/cms/review/:review_id", midd.Verify([]global.Role{global.AdminRole}), controller.DeleteCMSReview)
	v2.GET("/store/course/:course_id/reviews", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreCourseReviews)
	v2.GET("/store/course/review/:review_id", midd.Verify([]global.Role{global.UserRole}), controller.GetStoreCourseReview)
	v2.POST("/store/course/:course_id/review", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.CreateStoreCourseReview)
	v2.DELETE("/store/course/review/:review_id", middleware.Transaction(orm.Shared().DB()), midd.Verify([]global.Role{global.UserRole}), controller.DeleteStoreCourseReview)
}
