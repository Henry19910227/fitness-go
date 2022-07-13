package user

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user"

type Resolver interface {
	APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput)
	APIRegisterEmail(input *model.APIRegisterEmailInput) (output model.APIRegisterEmailOutput)
}
