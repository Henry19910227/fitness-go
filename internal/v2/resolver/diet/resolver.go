package diet

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	dietModel "github.com/Henry19910227/fitness-go/internal/v2/model/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet/api_create_diet"
	"github.com/Henry19910227/fitness-go/internal/v2/model/diet/api_get_diet"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	rdaModel "github.com/Henry19910227/fitness-go/internal/v2/model/rda"
	"github.com/Henry19910227/fitness-go/internal/v2/service/diet"
	"github.com/Henry19910227/fitness-go/internal/v2/service/rda"
)

type resolver struct {
	dietService diet.Service
	rdaService  rda.Service
}

func New(dietService diet.Service, rdaService rda.Service) Resolver {
	return &resolver{dietService: dietService, rdaService: rdaService}
}

func (r *resolver) APICreateDiet(input *api_create_diet.Input) (output api_create_diet.Output) {
	// 查詢最新 RDA
	rdaListInput := rdaModel.ListInput{}
	rdaListInput.UserID = util.PointerInt64(input.UserID)
	rdaListInput.OrderField = "create_at"
	rdaListInput.OrderType = orderByModel.DESC
	rdaListInput.Size = util.PointerInt(1)
	rdaOutputs, _, err := r.rdaService.List(&rdaListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(rdaOutputs) == 0 {
		output.Set(code.BadRequest, "尚未設定RDA")
		return output
	}
	// 創建 Diet
	dietTable := dietModel.Table{}
	dietTable.UserID = util.PointerInt64(input.UserID)
	dietTable.RdaID = rdaOutputs[0].ID
	dietTable.ScheduleAt = util.PointerString(input.Body.ScheduleAt)
	dietID, err := r.dietService.Create(&dietTable)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	findDietInput := dietModel.FindInput{}
	findDietInput.ID = util.PointerInt64(dietID)
	findDietInput.Preloads = []*preloadModel.Preload{
		{Field: "RDA"},
	}
	dietOutput, err := r.dietService.Find(&findDietInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := api_create_diet.Data{}
	if err := util.Parser(dietOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetDiet(input *api_get_diet.Input) (output api_get_diet.Output) {
	// 查詢 Diet
	dietListInput := dietModel.ListInput{}
	dietListInput.UserID = util.PointerInt64(input.UserID)
	dietListInput.ScheduleAt = util.PointerString(input.Query.ScheduleAt)
	dietListInput.Preloads = []*preloadModel.Preload{
		{Field: "RDA"},
		{Field: "Meals.Food.FoodCategory"},
	}
	dietOutputs, _, err := r.dietService.List(&dietListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(dietOutputs) == 0 {
		// Parse Output
		data := api_get_diet.Data{}
		data.ScheduleAt = util.PointerString(input.Query.ScheduleAt)
		// 查詢 RDA
		rdaListInput := rdaModel.ListInput{}
		rdaListInput.UserID = util.PointerInt64(input.UserID)
		rdaListInput.OrderField = "create_at"
		rdaListInput.OrderType = orderByModel.DESC
		rdaListInput.Size = util.PointerInt(1)
		rdaOutputs, _, err := r.rdaService.List(&rdaListInput)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		if len(rdaOutputs) > 0 {
			if err := util.Parser(rdaOutputs[0], &data.RDA); err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
		}
		output.Set(code.Success, "success")
		output.Data = &data
		return output
	}
	// Parse Output
	data := api_get_diet.Data{}
	if err := util.Parser(dietOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
