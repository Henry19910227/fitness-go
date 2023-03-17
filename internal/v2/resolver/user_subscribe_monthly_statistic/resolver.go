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
	listInput := model.ListInput{}
	listInput.Month = util.PointerInt(input.Query.Month)
	listInput.Year = util.PointerInt(input.Query.Year)
	statisticOutputs, _, err := r.statisticService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := api_get_cms_user_subscribe_statistic.Data{}
	data.Year = util.PointerInt(input.Query.Year)
	data.Month = util.PointerInt(input.Query.Month)
	data.Total = util.PointerInt(0)
	data.Male = util.PointerInt(0)
	data.Female = util.PointerInt(0)
	data.Age13to17 = util.PointerInt(0)
	data.Age18to24 = util.PointerInt(0)
	data.Age25to34 = util.PointerInt(0)
	data.Age35to44 = util.PointerInt(0)
	data.Age45to54 = util.PointerInt(0)
	data.Age55to64 = util.PointerInt(0)
	data.Age65Up = util.PointerInt(0)
	if len(statisticOutputs) > 0 {
		if err := util.Parser(statisticOutputs[0], &data); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
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
