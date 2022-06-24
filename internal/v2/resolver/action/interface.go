package action

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
)

type Resolver interface {
	APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput)
	APICreateCMSAction(input *model.APICreateCMSActionInput) (output model.APICreateCMSActionOutput)
	APIUpdateCMSAction(input *model.APIUpdateCMSActionInput) (output base.Output)
}
