package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_training_avg_statistic/api_get_cms_statistic_monthly_course_training_avg"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_training_avg_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver course_training_avg_statistic.Resolver
}

func New(resolver course_training_avg_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSCourseTrainingAvgStatistic 獲取課表平均訓練率列表
// @Summary 獲取課表平均訓練率列表
// @Description 獲取課表平均訓練率列表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param course_id query int false "課表id"
// @Param name query string false "課表名稱"
// @Param course_status query int false "課表狀態 (1:準備中/2:審核中/3:銷售中/4:退審/5:下架)"
// @Param sale_type query int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} api_get_cms_statistic_monthly_course_training_avg.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/course/training_avg [GET]
func (c *controller) GetCMSCourseTrainingAvgStatistic(ctx *gin.Context) {
	var input api_get_cms_statistic_monthly_course_training_avg.Input
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourseTrainingAvgStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}

func (c *controller) Statistic() {
	c.resolver.Statistic()
}
