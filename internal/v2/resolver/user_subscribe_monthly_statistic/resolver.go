package user_subscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_monthly_statistic/api_get_cms_user_subscribe_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_monthly_statistic"
	"strconv"
	"time"
)

type resolver struct {
	statisticService user_subscribe_monthly_statistic.Service
}

func New(statisticService user_subscribe_monthly_statistic.Service) Resolver {
	return &resolver{statisticService: statisticService}
}

func (r *resolver) APIGetCMSUserSubscribeStatistic(input *api_get_cms_user_subscribe_statistic.Input) (output api_get_cms_user_subscribe_statistic.Output) {
	currentYear, _ := strconv.Atoi(time.Now().Format("2006"))
	currentMonth, _ := strconv.Atoi(time.Now().Format("01"))
	if input.Query.Year > currentYear {
		output.Set(code.BadRequest, "不可大於當前時間")
		return output
	}
	if input.Query.Year == currentYear && input.Query.Month > currentMonth {
		output.Set(code.BadRequest, "不可大於當前時間")
		return output
	}
	// 查找是否有統計資料
	listInput := model.ListInput{}
	listInput.Month = util.PointerInt(input.Query.Month)
	listInput.Year = util.PointerInt(input.Query.Year)
	statisticOutputs, _, err := r.statisticService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := api_get_cms_user_subscribe_statistic.Data{}
	// 存在統計資料就 parser
	if len(statisticOutputs) > 0 {
		if err := util.Parser(statisticOutputs[0], &data); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	// 不存在統計資料就統計一次
	statisticInput := model.StatisticInput{}
	statisticInput.Year = input.Query.Year
	statisticInput.Month = input.Query.Month
	if err := r.statisticService.Statistic(&statisticInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查找
	findInput := model.FindInput{}
	findInput.Month = util.PointerInt(input.Query.Month)
	findInput.Year = util.PointerInt(input.Query.Year)
	statisticOutput, err := r.statisticService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parser
	if err := util.Parser(statisticOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) Statistic() {
	year, _ := strconv.Atoi(time.Now().Format("2006"))
	month, _ := strconv.Atoi(time.Now().Format("01"))
	statisticInput := model.StatisticInput{}
	statisticInput.Year = year
	statisticInput.Month = month
	_ = r.statisticService.Statistic(&statisticInput)
}
