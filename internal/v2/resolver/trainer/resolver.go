package trainer

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
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

func (r *resolver) APIGetStoreTrainer(input *model.APIGetStoreTrainerInput) (output model.APIGetStoreTrainerOutput) {
	findInput := model.FindInput{}
	findInput.UserID = util.PointerInt64(input.Uri.UserID)
	findInput.Preloads = []*preload.Preload{
		{Field: "User"},
		{Field: "TrainerStatistic"},
		{Field: "Certificates"},
		{Field: "TrainerAlbums"},
	}
	trainerOutput, err := r.trainerService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetStoreTrainerData{}
	data.IsDeleted = trainerOutput.UserOnSafe().IsDeleted
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreTrainers(input *model.APIGetStoreTrainersInput) (output model.APIGetStoreTrainersOutput) {
	joins := make([]*joinModel.Join, 0)
	wheres := make([]*whereModel.Where, 0)
	orders := make([]*orderByModel.Order, 0)
	if input.Query.OrderField != nil {
		if *input.Query.OrderField == "latest" {
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainers.%s %s", "create_at", order_by.DESC)})
		}
		if *input.Query.OrderField == "popular" {
			joins = append(joins, &joinModel.Join{Query: "LEFT JOIN trainer_statistics ON trainers.user_id = trainer_statistics.user_id"})
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainer_statistics.%s %s", "student_count", order_by.DESC)})
		}
	} else {
		orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainers.%s %s", "create_at", order_by.DESC)})
	}
	listInput := model.ListInput{}
	listInput.Wheres = wheres
	listInput.Joins = joins
	listInput.Orders = orders
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	trainerOutputs, page, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetStoreTrainersData{}
	if err := util.Parser(trainerOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
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
