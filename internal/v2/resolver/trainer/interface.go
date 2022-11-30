package trainer

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer/api_update_cms_trainer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateTrainer(tx *gorm.DB, input *model.APICreateTrainerInput) (output model.APICreateTrainerOutput)
	APIUpdateTrainer(tx *gorm.DB, input *model.APIUpdateTrainerInput) (output model.APIUpdateTrainerOutput)
	APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput)
	APIGetStoreTrainer(input *model.APIGetStoreTrainerInput) (output model.APIGetStoreTrainerOutput)
	APIGetStoreTrainers(input *model.APIGetStoreTrainersInput) (output model.APIGetStoreTrainersOutput)
	APIGetFavoriteTrainers(input *model.APIGetFavoriteTrainersInput) (output model.APIGetFavoriteTrainersOutput)

	APIUpdateCMSTrainer(ctx *gin.Context, tx *gorm.DB, input *api_update_cms_trainer.Input) (output api_update_cms_trainer.Output)
	APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput)
}
