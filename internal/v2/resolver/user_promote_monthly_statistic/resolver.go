package user_promote_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_promote_monthly_statistic"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user_promote_monthly_statistic/api_get_cms_statistic_monthly_user_promote"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_promote_monthly_statistic"
	"strconv"
	"time"
)

type resolver struct {
	statisticService user_promote_monthly_statistic.Service
}

func New(statisticService user_promote_monthly_statistic.Service) Resolver {
	return &resolver{statisticService: statisticService}
}

func (r *resolver) APIGetCMSUserPromoteStatistic(input *api_get_cms_statistic_monthly_user_promote.Input) (output api_get_cms_statistic_monthly_user_promote.Output) {
	data := api_get_cms_statistic_monthly_user_promote.Data{}
	data.Year = util.PointerInt(input.Query.Year)
	data.Month = util.PointerInt(input.Query.Month)
	data.Total = util.PointerInt(1000)
	data.Male = util.PointerInt(600)
	data.Female = util.PointerInt(400)
	data.Exp1to3 = util.PointerInt(100)
	data.Exp4to6 = util.PointerInt(100)
	data.Exp7to10 = util.PointerInt(200)
	data.Exp11to15 = util.PointerInt(200)
	data.Exp16to19 = util.PointerInt(200)
	data.Exp20up = util.PointerInt(200)
	data.Age13to17 = util.PointerInt(100)
	data.Age18to24 = util.PointerInt(150)
	data.Age25to34 = util.PointerInt(250)
	data.Age35to44 = util.PointerInt(200)
	data.Age45to54 = util.PointerInt(150)
	data.Age55to64 = util.PointerInt(100)
	data.Age65Up = util.PointerInt(50)
	data.CreateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	data.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
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
