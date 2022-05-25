package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/gin-gonic/gin"
)

type Statistic struct {
	Base
	incomeStatisticService service.UserIncomeMonthlyStatistic
}

func NewStatistic(baseGroup *gin.RouterGroup, incomeStatisticService service.UserIncomeMonthlyStatistic, userMidd midd.User) {
	statistic := Statistic{incomeStatisticService: incomeStatisticService}
	baseGroup.GET("/income_monthly_statistic",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		statistic.GetIncomeMonthlyStatistic)
}

// GetIncomeMonthlyStatistic 取得當月收益分析
// @Summary 取得當月收益分析
// @Description 取得當月收益分析
// @Tags Statistic
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.UserIncomeMonthlyStatistic} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /income_monthly_statistic [GET]
func (s *Statistic) GetIncomeMonthlyStatistic(c *gin.Context) {
	uid, e := s.GetUID(c)
	if e != nil {
		s.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	income, err := s.incomeStatisticService.GetUserIncomeMonthlyStatistic(c, uid)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, income, "success!")
}

// GetCourseUsageMonthlyStatistic 取得當月課表使用人數分析
// @Summary 取得當月課表使用人數分析
// @Description 取得當月課表使用人數分析
// @Tags Statistic
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.UserCourseUsageMonthlyStatistic} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /course_usage_monthly_statistic [GET]
func (s *Statistic) GetCourseUsageMonthlyStatistic(c *gin.Context) {
	//uid, e := t.GetUID(c)
	//if e != nil {
	//	t.JSONValidatorErrorResponse(c, e.Error())
	//	return
	//}
	//var uri validator.TrainerIDUri
	//if err := c.ShouldBindUri(&uri); err != nil {
	//	t.JSONValidatorErrorResponse(c, err.Error())
	//	return
	//}
	//trainer, err := t.trainerService.GetTrainer(c, &uid, uri.TrainerID)
	//if err != nil {
	//	t.JSONErrorResponse(c, err)
	//	return
	//}
	//t.JSONSuccessResponse(c, trainer, "success!")
}
