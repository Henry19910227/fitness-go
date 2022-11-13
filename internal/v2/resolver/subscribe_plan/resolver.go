package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan/api_get_subscribe_plans"
	"github.com/Henry19910227/fitness-go/internal/v2/service/subscribe_plan"
)

type resolver struct {
	subscribePlanService subscribe_plan.Service
}

func New(subscribePlanService subscribe_plan.Service) Resolver {
	return &resolver{subscribePlanService: subscribePlanService}
}

func (r *resolver) APIGetSubscribePlans(input *api_get_subscribe_plans.Input) (output api_get_subscribe_plans.Output) {
	// 查詢列表
	listInput := model.ListInput{}
	listInput.Enable = util.PointerInt(1)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "ProductLabel"},
	}
	subscribePlanOutputs, _, err := r.subscribePlanService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_subscribe_plans.Data{}
	if err := util.Parser(subscribePlanOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
