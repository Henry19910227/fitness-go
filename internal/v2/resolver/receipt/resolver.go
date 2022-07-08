package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	receiptService "github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
)

type resolver struct {
	receiptService receiptService.Service
}

func New(receiptService receiptService.Service) Resolver {
	return &resolver{receiptService: receiptService}
}

func (r *resolver) APIGetCMSOrderReceipts(input *model.APIGetCMSOrderReceiptsInput) (output model.APIGetCMSOrderReceiptsOutput) {
	// parser input
	param := model.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "ProductLabel"},
	}
	if err := util.Parser(input.Uri, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// get list
	datas, page, err := r.receiptService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSOrderReceiptsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}
