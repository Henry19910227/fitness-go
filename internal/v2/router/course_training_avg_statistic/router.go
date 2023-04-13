package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/scheduler"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/course_training_avg_statistic"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := course_training_avg_statistic.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.GET("/cms/statistic_monthly/course/training_avg", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSCourseTrainingAvgStatistic)

	_, _ = scheduler.Shared().Cron().AddFunc("0 0 * * * *", controller.Statistic)
}
