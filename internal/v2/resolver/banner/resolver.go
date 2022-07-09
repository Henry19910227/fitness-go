package banner

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	bannerService "github.com/Henry19910227/fitness-go/internal/v2/service/banner"
)

type resolver struct {
	bannerService bannerService.Service
	uploadTool   uploader.Tool
}

func New(bannerService bannerService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{bannerService: bannerService, uploadTool: uploadTool}
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

