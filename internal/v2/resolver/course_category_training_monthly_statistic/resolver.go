package course_category_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic"
	"time"
)

type resolver struct {
}

func New() Resolver {
	return &resolver{}
}

func (r *resolver) APIGetCMSCategoryTrainingStatistic(input *model.APIGetCMSCategoryTrainingStatisticInput) (output model.APIGetCMSCategoryTrainingStatisticOutput) {
	data := model.APIGetCMSCategoryTrainingStatisticData{}
	data.ID = util.PointerInt64(1)
	data.Category = util.PointerInt(input.Query.Category)
	data.Year = util.PointerInt(input.Query.Year)
	data.Month = util.PointerInt(input.Query.Month)
	data.Total = util.PointerInt(1000)
	data.Male = util.PointerInt(600)
	data.Female = util.PointerInt(400)
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
