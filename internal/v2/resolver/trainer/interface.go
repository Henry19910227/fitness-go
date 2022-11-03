package trainer

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateTrainer(tx *gorm.DB, input *model.APICreateTrainerInput) (output model.APICreateTrainerOutput)
	APIUpdateTrainer(tx *gorm.DB, input *model.APIUpdateTrainerInput) (output model.APIUpdateTrainerOutput)
	APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput)
	APIGetStoreTrainer(input *model.APIGetStoreTrainerInput) (output model.APIGetStoreTrainerOutput)
	APIGetStoreTrainers(input *model.APIGetStoreTrainersInput) (output model.APIGetStoreTrainersOutput)
	APIGetFavoriteTrainers(input *model.APIGetFavoriteTrainersInput) (output model.APIGetFavoriteTrainersOutput)
	APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput)
}
