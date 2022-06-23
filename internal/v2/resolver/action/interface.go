package action

import model "github.com/Henry19910227/fitness-go/internal/v2/model/action"

type Resolver interface {
	APIGetCMSActions(input *model.APIGetCMSActionsInput) (output model.APIGetCMSActionsOutput)
}
