package user_subscribe_info

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
)

type resolver struct {
	subscribeInfoService user_subscribe_info.Service
}

func New(subscribeInfoService user_subscribe_info.Service) Resolver {
	return &resolver{subscribeInfoService: subscribeInfoService}
}

func (r *resolver) APIGetUserSubscribeInfo(input *model.APIGetUserSubscribeInfoInput) (output model.APIGetUserSubscribeInfoOutput) {
	findInput := model.FindInput{}
	findInput.UserID = util.PointerInt64(input.UserID)
	outputData, err := r.subscribeInfoService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetUserSubscribeInfoData{}
	if err := util.Parser(outputData, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}