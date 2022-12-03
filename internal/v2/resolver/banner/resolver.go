package banner

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
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
	listInput.Joins = []*joinModel.Join{
		{Query: "LEFT JOIN banner_orders ON banners.id = banner_orders.banner_id"},
	}
	listInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("banner_orders.seq IS NULL ASC, banner_orders.seq ASC, banners.create_at ASC")},
	}
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
	// 查詢 banner
	listInput := model.ListInput{}
	listInput.Joins = []*joinModel.Join{
		{Query: "LEFT JOIN banner_orders ON banners.id = banner_orders.banner_id"},
	}
	if input.Form.OrderField == "create_at" {
		listInput.OrderField = input.Form.OrderField
		listInput.OrderType = input.Form.OrderType
	}
	if input.Form.OrderField == "seq" {
		listInput.Orders = []*orderByModel.Order{
			{Value: fmt.Sprintf("banner_orders.seq IS NULL %v, banner_orders.seq %v, banners.create_at %v", input.Form.OrderType, input.Form.OrderType, input.Form.OrderType)},
		}
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "Course"},
	}
	listInput.Page = input.Form.Page
	listInput.Size = input.Form.Size
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
