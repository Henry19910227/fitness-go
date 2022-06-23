package trainer

import model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"

type Resolver interface {
	APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput)
}