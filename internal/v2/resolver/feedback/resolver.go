package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/feedback"
	imageModel "github.com/Henry19910227/fitness-go/internal/v2/model/feedback_image"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	"github.com/Henry19910227/fitness-go/internal/v2/service/feedback"
	feedbackImage "github.com/Henry19910227/fitness-go/internal/v2/service/feedback_image"
	"gorm.io/gorm"
)

type resolver struct {
	feedbackService      feedback.Service
	feedbackImageService feedbackImage.Service
	uploadTool           uploader.Tool
}

func New(feedbackService feedback.Service, feedbackImageService feedbackImage.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{feedbackService: feedbackService, feedbackImageService: feedbackImageService, uploadTool: uploadTool}
}

func (r *resolver) APICreateFeedback(tx *gorm.DB, input *model.APICreateFeedbackInput) (output base.Output) {
	defer tx.Rollback()
	//創建feedback
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	if err := util.Parser(input.Form, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	feedbackID, err := r.feedbackService.Tx(tx).Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//創建feedback_image
	imageTables := make([]*imageModel.Table, 0)
	for _, file := range input.Files {
		imageNamed, err := r.uploadTool.Save(file.Data, file.Named)
		if err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		table := imageModel.Table{}
		table.FeedbackID = util.PointerInt64(feedbackID)
		table.Image = util.PointerString(imageNamed)
		imageTables = append(imageTables, &table)
	}
	if err := r.feedbackImageService.Tx(tx).Create(imageTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIGetCMSFeedbacks(input *model.APIGetCMSFeedbacksInput) (output model.APIGetCMSFeedbacksOutput) {
	// parser input
	param := model.ListInput{}
	param.Preloads = []*preloadModel.Preload{
		{Field: "Images"},
		{Field: "User"},
	}
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 調用 service
	result, page, err := r.feedbackService.List(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetCMSFeedbacksData{}
	if err := util.Parser(result, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = data
	output.Paging = page
	return output
}
