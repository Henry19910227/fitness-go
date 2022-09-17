package course_category_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_category_training_monthly_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver course_category_training_monthly_statistic.Resolver
}

func New(resolver course_category_training_monthly_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSCategoryTrainingMonthlyStatistic 獲取課表分類訓練統計月報表
// @Summary 獲取課表分類訓練統計月報表
// @Description 獲取課表分類訓練統計月報表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param category query int true "課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)"
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} course_category_training_monthly_statistic.APIGetCMSCategoryTrainingStatisticOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/course_category/training [GET]
func (c *controller) GetCMSCategoryTrainingMonthlyStatistic(ctx *gin.Context) {
	var input model.APIGetCMSCategoryTrainingStatisticInput
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCategoryTrainingStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}
