package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/orm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/scheduler"
	"github.com/Henry19910227/fitness-go/internal/v2/controller/trainer_statistic"
	"github.com/gin-gonic/gin"
)

func SetRoute(v2 *gin.RouterGroup) {
	controller := trainer_statistic.NewController(orm.Shared().DB())

	_, _ = scheduler.Shared().Cron().AddFunc("0 0 * * * *", controller.StatisticStudentCount)
}
