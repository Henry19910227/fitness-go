package food

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	foodService "github.com/Henry19910227/fitness-go/internal/v2/service/food"
	"github.com/gin-gonic/gin"
)

type resolver struct {
	foodService foodService.Service
	logTool     logger.Tool
}

func New(foodService foodService.Service, logTool logger.Tool) Resolver {
	return &resolver{foodService: foodService, logTool: logTool}
}

func (r *resolver) APIGetFoods(ctx *gin.Context, input *model.APIGetFoodsInput) interface{} {
	// parser input
	param := model.ListInput{}
	if err := util.Parser(input, &param); err != nil {
		r.logTool.Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	param.Preloads = []*preloadModel.Preload{
		{Field: "FoodCategory"},
	}
	// 調用 repo
	result, _, err := r.foodService.List(&param)
	if err != nil {
		r.logTool.Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	// parser output
	data := model.APIGetFoodsData{}
	if err := util.Parser(result, &data); err != nil {
		r.logTool.Error(ctx, err.Error())
		return base.BadRequest(util.PointerString(err.Error()))
	}
	output := &model.APIGetFoodsOutput{}
	output.Data = data
	output.Code = code.Success
	output.Msg = "success!"
	return output
}
