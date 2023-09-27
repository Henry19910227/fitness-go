package banner

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_create_cms_banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_delete_cms_banner"
	"github.com/Henry19910227/fitness-go/internal/v2/model/banner/api_get_cms_banners"
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

func (r *resolver) APICreateCMSBanner(input *api_create_cms_banner.Input) (output api_create_cms_banner.Output) {
	// 驗證格式
	if input.Form.Type == model.CourseType && input.Form.CourseID == nil {
		output.Set(code.BadRequest, "新增課表類型banner, course_id 欄位不得為空")
		return output
	}
	if input.Form.Type == model.TrainerType && input.Form.UserID == nil {
		output.Set(code.BadRequest, "新增教練類型banner, user_id 欄位不得為空")
		return output
	}
	if input.Form.Type == model.UrlType && input.Form.Url == nil {
		output.Set(code.BadRequest, "新增URL類型banner, url 欄位不得為空")
		return output
	}
	// parser table
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
	data := api_create_cms_banner.Data{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetCMSBanners(input *api_get_cms_banners.Input) (output api_get_cms_banners.Output) {
	// 查詢 banner
	listInput := model.ListInput{}
	listInput.Joins = []*joinModel.Join{
		{Query: "LEFT JOIN banner_orders ON banners.id = banner_orders.banner_id"},
	}
	if input.Query.OrderField == "create_at" {
		listInput.OrderField = input.Query.OrderField
		listInput.OrderType = input.Query.OrderType
	}
	if input.Query.OrderField == "seq" {
		listInput.Orders = []*orderByModel.Order{
			{Value: fmt.Sprintf("banner_orders.seq IS NULL %v, banner_orders.seq %v, banners.create_at %v", input.Query.OrderType, input.Query.OrderType, input.Query.OrderType)},
		}
	}
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "Course"},
	}
	listInput.Page = input.Query.Page
	listInput.Size = input.Query.Size
	datas, page, err := r.bannerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := api_get_cms_banners.Data{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIDeleteCMSBanner(input *api_delete_cms_banner.Input) (output api_delete_cms_banner.Output) {
	//查找banner
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.Uri.BannerID)
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
	deleteInput.ID = util.PointerInt64(input.Uri.BannerID)
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
