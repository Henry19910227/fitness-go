package user

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user"

type Resolver interface {
	APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput)
	APIRegisterForEmail(input *model.APIRegisterForEmailInput) (output model.APIRegisterForEmailOutput)
	APIRegisterForFacebook(input *model.APIRegisterForFacebookInput) (output model.APIRegisterForFacebookOutput)
	APILoginForEmail(input *model.APILoginForEmailInput) (output model.APILoginForEmailOutput)
	APILogout(input *model.APILogoutInput) (output model.APILogoutOutput)
	APICreateRegisterOTP(input *model.APICreateOTPInput) (output model.APICreateRegisterOTPOutput)
	APIRegisterNicknameValidate(input *model.APIRegisterNicknameValidateInput) (output model.APIRegisterNicknameValidateOutput)
	APIRegisterEmailAccountValidate(input *model.APIRegisterEmailAccountValidateInput) (output model.APIRegisterEmailAccountValidateOutput)
	APIRegisterFacebookAccountValidate(input *model.APIRegisterFacebookAccountValidateInput) (output model.APIRegisterFacebookAccountValidateOutput)
}
