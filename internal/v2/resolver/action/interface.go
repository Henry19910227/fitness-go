package action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_create_trainer_action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_get_trainer_course_actions"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_get_user_action_best_pr"
	"github.com/Henry19910227/fitness-go/internal/v2/model/action/api_update_trainer_action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserAction(tx *gorm.DB, input *model.APICreateUserActionInput) (output model.APICreateUserActionOutput)
	APIUpdateUserAction(tx *gorm.DB, input *model.APIUpdateUserActionInput) (output model.APIUpdateUserActionOutput)
	APIGetUserActions(input *model.APIGetUserActionsInput) (output model.APIGetUserActionsOutput)
	APIGetUserActionBestPR(input *api_get_user_action_best_pr.Input) (output api_get_user_action_best_pr.Output)
	APIDeleteUserAction(input *model.APIDeleteUserActionInput) (output model.APIDeleteUserActionOutput)
	APIDeleteUserActionVideo(input *model.APIDeleteUserActionVideoInput) (output model.APIDeleteUserActionVideoOutput)
	APIGetUserActionSystemImages() (output model.APIGetUserActionSystemImagesOutput)

	APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput)
	APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput)
	APIUpdateCMSAction(input *model.APIUpdateCMSActionInput) (output base.Output)

	APICreateTrainerAction(tx *gorm.DB, input *api_create_trainer_action.Input) (output api_create_trainer_action.Output)
	APIUpdateTrainerAction(tx *gorm.DB, input *api_update_trainer_action.Input) (output api_update_trainer_action.Output)
	APIGetTrainerCourseActions(input *api_get_trainer_course_actions.Input) (output api_get_trainer_course_actions.Output)
	APIDeleteTrainerAction(input *model.APIDeleteTrainerActionInput) (output model.APIDeleteTrainerActionOutput)
	APIDeleteTrainerActionVideo(input *model.APIDeleteTrainerActionVideoInput) (output model.APIDeleteTrainerActionVideoOutput)
}
