package action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"gorm.io/gorm"
)

type Resolver interface {
	APICreateUserAction(tx *gorm.DB, input *model.APICreateUserActionInput) (output model.APICreateUserActionOutput)
	APIUpdateUserAction(tx *gorm.DB, input *model.APIUpdateUserActionInput) (output model.APIUpdateUserActionOutput)
	APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput)
	APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput)
	APIUpdateCMSAction(input *model.APIUpdateCMSActionInput) (output base.Output)
}
