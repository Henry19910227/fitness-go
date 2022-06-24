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
	actionService   actionService.Service
	coverUploadTool uploader.Tool
	videoUploadTool uploader.Tool
}

func New(courseService actionService.Service, coverUploadTool uploader.Tool, videoUploadTool uploader.Tool) Resolver {
	return &resolver{actionService: courseService, coverUploadTool: coverUploadTool, videoUploadTool: videoUploadTool}
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

func (r *resolver) APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput) {
	table := model.Table{}
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table.Cover = util.PointerString("")
	table.Video = util.PointerString("")
	table.Source = util.PointerInt(1)
	// 儲存動作封面圖
	if input.CoverFile != nil {
		coverNamed, err := r.coverUploadTool.Save(input.CoverFile.Data, input.CoverFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Cover = util.PointerString(coverNamed)
	}
	// 儲存動作影片
	if input.VideoFile != nil {
		videoNamed, err := r.videoUploadTool.Save(input.VideoFile.Data, input.VideoFile.Named)
		if err != nil {
			if table.Cover != nil {
				_ = r.coverUploadTool.Delete(*table.Cover)
			}
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Video = util.PointerString(videoNamed)
	}
	result, err := r.actionService.Create(&table)
	if err != nil {
		if table.Cover != nil {
			_ = r.coverUploadTool.Delete(*table.Cover)
		}
		if table.Video != nil {
			_ = r.videoUploadTool.Delete(*table.Video)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateCMSActionData{}
	if err := util.Parser(result, &data); err != nil {
		if table.Cover != nil {
			_ = r.coverUploadTool.Delete(*table.Cover)
		}
		if table.Video != nil {
			_ = r.videoUploadTool.Delete(*table.Video)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) clean() {

}
