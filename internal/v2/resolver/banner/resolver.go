package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	bannerService "github.com/Henry19910227/fitness-go/internal/v2/service/banner"
)

type resolver struct {
	bannerService bannerService.Service
	uploadTool    uploader.Tool
}

func New(bannerService bannerService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{bannerService: bannerService, uploadTool: uploadTool}
}

func (r *resolver) APIGetBanners(input *model.APIGetBannersInput) (output model.APIGetBannersOutput) {
	// parser input
	listInput := model.ListInput{}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "Course"},
	}
	if err := util.Parser(input.Query, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// List
	datas, page, err := r.bannerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetBannersData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APICreateCMSBanner(input *model.APICreateCMSBannerInput) (output model.APICreateCMSBannerOutput) {
	table := model.Table{}
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存圖片
	if input.ImageFile != nil {
		imageNamed, err := r.uploadTool.Save(input.ImageFile.Data, input.ImageFile.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table.Image = util.PointerString(imageNamed)
	}
	// Create
	result, err := r.bannerService.Create(&table)
	if err != nil {
		if table.Image != nil {
			_ = r.uploadTool.Delete(*table.Image)
		}
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateCMSBannerData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetCMSBanners(input *model.APIGetCMSBannersInput) (output model.APIGetCMSBannersOutput) {
	// parser input
	listInput := model.ListInput{}
	if err := util.Parser(input.Form, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// List
	datas, page, err := r.bannerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSBannersData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIDeleteCMSBanner(input *model.APIDeleteCMSBannerInput) (output model.APIDeleteCMSBannerOutput) {
	//查找banner
	findInput := model.FindInput{}
	if err := util.Parser(input.Uri, &findInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	findOutput, err := r.bannerService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser delete input
	deleteInput := model.DeleteInput{}
	if err := util.Parser(input.Uri, &deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除banner
	if err := r.bannerService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除banner圖片
	_ = r.uploadTool.Delete(util.OnNilJustReturnString(findOutput.Image, ""))
	output.Set(code.Success, "success")
	return output
}
