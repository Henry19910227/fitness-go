package user

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user"

type Resolver interface {
	APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput)
	APIRegisterForEmail(input *model.APIRegisterForEmailInput) (output model.APIRegisterForEmailOutput)
	APILoginForEmail(input *model.APILoginForEmailInput) (output model.APILoginForEmailOutput)
	APICreateRegisterOTP(input *model.APICreateRegisterOTPInput) (output model.APICreateRegisterOTPOutput)
	APIRegisterNicknameValidate(input *model.APIRegisterNicknameValidateInput) (output model.APIRegisterNicknameValidateOutput)
	APIRegisterAccountValidate(input *model.APIRegisterAccountValidateInput) (output model.APIRegisterAccountValidateOutput)
}
