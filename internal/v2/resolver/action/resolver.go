package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	actionService "github.com/Henry19910227/fitness-go/internal/v2/service/action"
)

type resolver struct {
	actionService actionService.Service
	uploadTool    uploader.Tool
}

func New(courseService actionService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{actionService: courseService, uploadTool: uploadTool}
}

func (r *resolver) APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput) {
	actionInput := model.ListInput{}
	actionInput.Source = util.PointerInt(1)
	actionInput.Size = input.Size
	actionInput.Page = input.Page
	actionInput.OrderField = "create_at"
	actionInput.OrderType = order_by.ASC
	datas, page, err := r.actionService.List(&actionInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSActionsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}
