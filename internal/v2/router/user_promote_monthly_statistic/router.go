package user_promote_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/scheduler"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/user_promote_monthly_statistic"
	tokenMiddleware "github.com/Henry19910227/fitness-go/internal/v2/middleware/token"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := user_promote_monthly_statistic.NewController(orm.Shared().DB())
	midd := tokenMiddleware.NewTokenMiddleware(redis.Shared())
	v2.GET("/cms/statistic_monthly/user/promote", midd.Verify([]global.Role{global.AdminRole}), controller.GetCMSUserPromoteMonthlyStatistic)

	_, _ = scheduler.Shared().Cron().AddFunc("0 * * * * *", controller.Statistic) // 每小時統計一次
}
