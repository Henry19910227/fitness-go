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
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	subscribeInfoOutputs, _, err := r.subscribeInfoService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetUserSubscribeInfoData{}
	data.Status = util.PointerInt(0)
	if len(subscribeInfoOutputs) > 0 {
		if err := util.Parser(subscribeInfoOutputs[0], &data); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}