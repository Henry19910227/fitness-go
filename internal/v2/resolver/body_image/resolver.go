package body_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/body_image"
	bodyModel "github.com/Henry19910227/fitness-go/internal/v2/model/body_record"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	bodyImageService "github.com/Henry19910227/fitness-go/internal/v2/service/body_image"
	bodyService "github.com/Henry19910227/fitness-go/internal/v2/service/body_record"
)

type resolver struct {
	bodyImageService bodyImageService.Service
	bodyService      bodyService.Service
	bodyUploadTool   uploader.Tool
}

func New(bodyImageService bodyImageService.Service, bodyService bodyService.Service, bodyUploadTool uploader.Tool) Resolver {
	return &resolver{bodyImageService: bodyImageService, bodyUploadTool: bodyUploadTool, bodyService: bodyService}
}

func (r *resolver) APIGetBodyImages(input *model.APIGetBodyImagesInput) (output model.APIGetBodyImagesOutput) {
	// parser input
	bodyInput := model.ListInput{}
	bodyInput.UserID = util.PointerInt64(input.UserID)
	bodyInput.OrderField = "create_at"
	bodyInput.OrderType = order_by.DESC
	if err := util.Parser(input.Query, &bodyInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// List
	datas, page, err := r.bodyImageService.List(&bodyInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetBodyImagesData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APICreateBodyImage(input *model.APICreateBodyImageInput) (output model.APICreateBodyImageOutput) {
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.BodyImage = util.PointerString("")
	table.Weight = util.PointerFloat64(0)
	// 查找體重
	bodyListInput := bodyModel.ListInput{}
	bodyListInput.UserID = util.PointerInt64(input.UserID)
	bodyListInput.RecordType = util.PointerInt(1)
	bodyListInput.Page = 1
	bodyListInput.Size = 1
	bodyListInput.OrderField = "create_at"
	bodyListInput.OrderType = order_by.DESC
	records, _, err := r.bodyService.List(&bodyListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(records) > 0 {
		table.Weight = records[0].Value
	}
	// 儲存體態照片
	if input.ImageFile != nil {
		imageNamed, err := r.bodyUploadTool.Save(input.ImageFile.Data, input.ImageFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.BodyImage = util.PointerString(imageNamed)
	}
	// 創建體態照片
	result, err := r.bodyImageService.Create(&table)
	if err != nil {
		if table.BodyImage != nil {
			_ = r.bodyUploadTool.Delete(*table.BodyImage)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateBodyImageData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}
