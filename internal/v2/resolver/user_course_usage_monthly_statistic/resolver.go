package user_course_usage_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_usage_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_course_usage_monthly_statistic/api_get_trainer_course_usage_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_course_usage_monthly_statistic"
	"strconv"
	"time"
)

type resolver struct {
	statisticService user_course_usage_monthly_statistic.Service
}

func New(statisticService user_course_usage_monthly_statistic.Service) Resolver {
	return &resolver{statisticService: statisticService}
}

func (r *resolver) APIGetTrainerCourseUsageMonthlyStatistic(input *api_get_trainer_course_usage_monthly_statistic.Input) (output api_get_trainer_course_usage_monthly_statistic.Output) {
	// 獲取當前年份
	year, err := strconv.Atoi(time.Now().Format("2006"))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 獲取當前月份
	month, err := strconv.Atoi(time.Now().Format("1"))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢數據
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Year = util.PointerInt(year)
	listInput.Month = util.PointerInt(month)
	statisticOutputs, _, err := r.statisticService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	data := api_get_trainer_course_usage_monthly_statistic.Data{}
	data.Year = util.PointerInt(year)
	data.Month = util.PointerInt(month)
	if len(statisticOutputs) == 0 {
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	if err := util.Parser(statisticOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) Statistic() {
	_ = r.statisticService.Statistic()
}
