package user_subscribe_info

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"

type Resolver interface {
	APIGetUserSubscribeInfo(input *model.APIGetUserSubscribeInfoInput) (output model.APIGetUserSubscribeInfoOutput)
}
