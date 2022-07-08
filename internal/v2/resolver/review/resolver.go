package review

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	reviewService "github.com/Henry19910227/fitness-go/internal/v2/service/review"
)

type resolver struct {
	reviewService reviewService.Service
	uploadTool    uploader.Tool
}

func New(orderService reviewService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{reviewService: orderService, uploadTool: uploadTool}
}

func (r *resolver) APIGetCMSReviews(input *model.APIGetCMSReviewsInput) (output model.APIGetCMSReviewsOutput) {
	// parser input
	param := model.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Course"},
		{Field: "User"},
		{Field: "Images"},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// get list
	datas, page, err := r.reviewService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSReviewsData{}
	if err := util.Parser(datas, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIUpdateCMSReview(input *model.APIUpdateCMSReviewInput) (output model.APIUpdateCMSReviewOutput) {
	//parser input
	table := model.Table{}
	if err := util.Parser(input.Uri, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//執行更新
	if err := r.reviewService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIDeleteCMSReview(input *model.APIDeleteCMSReviewInput) (output model.APIDeleteCMSReviewOutput) {
	//查找評論
	findInput := model.FindInput{}
	if err := util.Parser(input.Uri, &findInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Images"},
	}
	findOutput, err := r.reviewService.Find(&findInput)
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
	//刪除評論
	if err := r.reviewService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除評論照片
	for _, imageOutput := range findOutput.Images {
		_ = r.uploadTool.Delete(util.OnNilJustReturnString(imageOutput.Image, ""))
	}
	output.Set(code.Success, "success")
	return output
}
