package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	trainerService "github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
)

type resolver struct {
	trainerService trainerService.Service
	uploadTool     uploader.Tool
}

func New(trainerService trainerService.Service, uploadTool uploader.Tool) Resolver {
	return &resolver{trainerService: trainerService, uploadTool: uploadTool}
}

func (r *resolver) APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Preloads = []*preload.Preload{{Field: "TrainerStatistic"}, {Field: "Certificates"}, {Field: "TrainerAlbums"}}
	listInput.Size = 1
	listInput.Page = 1
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	trainerOutputs, _, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(trainerOutputs) == 0 {
		output.Set(code.DataNotFound, "查無資料")
		return output
	}
	data := model.APIGetTrainerProfileData{}
	if err := util.Parser(trainerOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetFavoriteTrainers(input *model.APIGetFavoriteTrainersInput) (output model.APIGetFavoriteTrainersOutput) {
	// parser input
	param := model.FavoriteListInput{}
	param.UserID = util.PointerInt64(input.UserID)
	param.OrderField = "create_at"
	param.OrderType = order_by.DESC
	if err := util.Parser(input.Form, &param); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 執行查詢
	results, page, err := r.trainerService.FavoriteList(&param)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetFavoriteTrainersData{}
	if err := util.Parser(results, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.Avatar = util.PointerString(fileNamed)
	if err := r.trainerService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}
