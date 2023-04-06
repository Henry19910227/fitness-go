package course_create_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course_create_monthly_statistic/api_get_cms_statistic_monthly_course_create"
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

// GetCMSStatisticMonthlyCourseCreate 獲取課表創建統計月報表
// @Summary 獲取課表創建統計月報表
// @Description 獲取課表創建統計月報表
// @Tags CMS數據管理_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param year query int true "年份"
// @Param month query int true "月份"
// @Success 200 {object} api_get_cms_statistic_monthly_course_create.Output "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/statistic_monthly/course/create [GET]
func (c *controller) GetCMSStatisticMonthlyCourseCreate(ctx *gin.Context) {
	var input api_get_cms_statistic_monthly_course_create.Input
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetCMSStatisticMonthlyCourseCreate(&input)
	ctx.JSON(http.StatusOK, output)
}

func (c *controller) Statistic() {
	c.resolver.Statistic()
}
