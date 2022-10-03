package action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserAction(tx *gorm.DB, input *model.APICreateUserActionInput) (output model.APICreateUserActionOutput)
	APIUpdateUserAction(tx *gorm.DB, input *model.APIUpdateUserActionInput) (output model.APIUpdateUserActionOutput)
	APIGetUserActions(input *model.APIGetUserActionsInput) (output model.APIGetUserActionsOutput)
	APIDeleteUserAction(input *model.APIDeleteUserActionInput) (output model.APIDeleteUserActionOutput)
	APIDeleteUserActionVideo(input *model.APIDeleteUserActionVideoInput) (output model.APIDeleteUserActionVideoOutput)
	APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput)
	APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput)
	APIUpdateCMSAction(input *model.APIUpdateCMSActionInput) (output base.Output)

	APICreateTrainerAction(tx *gorm.DB, input *model.APICreateTrainerActionInput) (output model.APICreateTrainerActionOutput)
	APIUpdateTrainerAction(tx *gorm.DB, input *model.APIUpdateTrainerActionInput) (output model.APIUpdateTrainerActionOutput)
	APIGetTrainerActions(input *model.APIGetTrainerActionsInput) (output model.APIGetTrainerActionsOutput)
	APIDeleteTrainerAction(input *model.APIDeleteTrainerActionInput) (output model.APIDeleteTrainerActionOutput)
	APIDeleteTrainerActionVideo(input *model.APIDeleteTrainerActionVideoInput) (output model.APIDeleteTrainerActionVideoOutput)
}
