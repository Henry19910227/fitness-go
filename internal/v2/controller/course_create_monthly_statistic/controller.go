package course_create_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_create_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_create_monthly_statistic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver course_create_monthly_statistic.Resolver
}

func New(resolver course_create_monthly_statistic.Resolver) Controller {
	return &controller{resolver: resolver}
}

// GetCMSCourseCreateMonthlyStatistic 獲取課表創建統計月報表
// @Summary 獲取課表創建統計月報表
// @Description 獲取課表創建統計月報表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} course_create_monthly_statistic.APIGetCMSCourseCreateStatisticOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/course/create [GET]
func (c *controller) GetCMSCourseCreateMonthlyStatistic(ctx *gin.Context) {
	var input model.APIGetCMSCourseCreateStatisticInput
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSCourseCreateStatistic(&input)
	ctx.JSON(http.StatusOK, output)
}
