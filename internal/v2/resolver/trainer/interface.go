package trainer

import model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"

type Resolver interface {
	APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput)
	APIGetTrainer(input *model.APIGetTrainerInput) (output model.APIGetTrainerOutput)
	APIGetFavoriteTrainers(input *model.APIGetFavoriteTrainersInput) (output model.APIGetFavoriteTrainersOutput)
	APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput)
}
