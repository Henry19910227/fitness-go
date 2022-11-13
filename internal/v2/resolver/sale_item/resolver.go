package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/sale_item/api_get_sale_items"
	"github.com/Henry19910227/fitness-go/internal/v2/service/sale_item"
)

type resolver struct {
	saleItemService sale_item.Service
}

func New(saleItemService sale_item.Service) Resolver {
	return &resolver{saleItemService: saleItemService}
}

func (r *resolver) APIGetSaleItems(input *api_get_sale_items.Input) (output api_get_sale_items.Output) {
	// 查詢列表
	listInput := model.ListInput{}
	listInput.Type = util.PointerInt(model.ChargeCourse)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "ProductLabel"},
	}
	saleItemOutputs, _, err := r.saleItemService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_sale_items.Data{}
	if err := util.Parser(saleItemOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
