package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review_image"
	reviewImageService "github.com/Henry19910227/fitness-go/internal/v2/service/review_image"
)

type resolver struct {
	reviewImageService reviewImageService.Service
	uploadTool         uploader.Tool
}

func New(reviewImageService reviewImageService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{reviewImageService: reviewImageService, uploadTool: uploadTool}
}

func (r *resolver) APIDeleteCMSReviewImage(input *model.APIDeleteCMSReviewImageInput) (output model.APIDeleteCMSReviewImageOutput) {
	//查找評論照片
	findInput := model.FindInput{}
	if err := util.Parser(input.Uri, &findInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	findOutput, err := r.reviewImageService.Find(&findInput)
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
	if err := r.reviewImageService.Delete(&deleteInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//刪除評論照片
	_ = r.uploadTool.Delete(util.OnNilJustReturnString(findOutput.Image, ""))
	output.Set(code.Success, "success")
	return output
}
