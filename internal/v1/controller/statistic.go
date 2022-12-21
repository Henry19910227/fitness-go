package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	midd "github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/gin-gonic/gin"
)

type Statistic struct {
	Base
	incomeStatisticService service.UserIncomeMonthlyStatistic
	usageStatisticService  service.UserCourseUsageMonthlyStatistic
}

func NewStatistic(baseGroup *gin.RouterGroup, incomeStatisticService service.UserIncomeMonthlyStatistic, usageStatisticService service.UserCourseUsageMonthlyStatistic, userMidd midd.User) {
	statistic := Statistic{incomeStatisticService: incomeStatisticService, usageStatisticService: usageStatisticService}
	baseGroup.GET("/income_monthly_statistic",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		statistic.GetIncomeMonthlyStatistic)

	baseGroup.GET("/course_usage_monthly_statistic",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		statistic.GetCourseUsageMonthlyStatistic)
}

// GetIncomeMonthlyStatistic 取得當月收益分析
// @Summary 取得當月收益分析 (API已過時，更新為 /v2/trainer/income_monthly_statistic [GET])
// @Description 取得當月收益分析
// @Tags Statistic_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.UserIncomeMonthlyStatistic} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/income_monthly_statistic [GET]
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
// @Summary 取得當月課表使用人數分析 (API已過時，更新為 /v2/trainer/course/usage_monthly_statistic [GET])
// @Description 取得當月課表使用人數分析
// @Tags Statistic_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.UserCourseUsageMonthlyStatistic} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/course_usage_monthly_statistic [GET]
func (s *Statistic) GetCourseUsageMonthlyStatistic(c *gin.Context) {
	uid, e := s.GetUID(c)
	if e != nil {
		s.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	result, err := s.usageStatisticService.GetUserCourseUsageMonthlyStatistic(c, uid)
	if err != nil {
		s.JSONErrorResponse(c, err)
		return
	}
	s.JSONSuccessResponse(c, result, "success!")
}
