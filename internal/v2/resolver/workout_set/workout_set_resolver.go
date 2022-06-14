package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	workoutSetService "github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
)

type resolver struct {
	workoutSetService workoutSetService.Service
}

func New(workoutSetService workoutSetService.Service) Resolver {
	return &resolver{workoutSetService: workoutSetService}
}

func (r *resolver) APIGetCMSWorkoutSets(input *model.APIGetCMSWorkoutSetsInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Action"},
	}
	// 調用 repo
	result, page, err := r.workoutSetService.List(&param)
	if err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetCMSWorkoutSetsData{}
	if err := util.Parser(result, &data); err != nil {
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetCMSWorkoutSetsOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	output.Paging = page
	return output
}
