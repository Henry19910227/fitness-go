package plan

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	planService "github.com/Henry19910227/fitness-go/internal/v2/service/plan"
)

type resolver struct {
	planService planService.Service
}

func New(planService planService.Service) Resolver {
	return &resolver{planService: planService}
}

func (r *resolver) APIGetCMSPlans(input *model.APIGetCMSPlansInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Workout", OrderBy: order_by.NewInput("create_at", "DESC")},
	}
	// 調用 repo
	result, page, err := r.planService.List(&param)
	if err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSPlansData{}
	if err := util.Parser(result, &data); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSPlansOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}
