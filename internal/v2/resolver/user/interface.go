package user

import model "github.com/Henry19910227/fitness-go/internal/v2/model/user"

type Resolver interface {
	APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput)
	APIUpdateUserProfile(input *model.APIUpdateUserProfileInput) (output model.APIUpdateUserProfileOutput)
	APIUpdateUserAvatar(input *model.APIUpdateUserAvatarInput) (output model.APIUpdateUserAvatarOutput)
	APIGetUserProfile(input *model.APIGetUserProfileInput) (output model.APIGetUserProfileOutput)
	APIRegisterForEmail(input *model.APIRegisterForEmailInput) (output model.APIRegisterForEmailOutput)
	APIRegisterForFacebook(input *model.APIRegisterForFacebookInput) (output model.APIRegisterForFacebookOutput)
	APIRegisterForGoogle(input *model.APIRegisterForGoogleInput) (output model.APIRegisterForGoogleOutput)
	APIRegisterForLine(input *model.APIRegisterForLineInput) (output model.APIRegisterForLineOutput)
	APILoginForEmail(input *model.APILoginForEmailInput) (output model.APILoginForEmailOutput)
	APILoginForFacebook(input *model.APILoginForFacebookInput) (output model.APILoginForFacebookOutput)
	APILoginForGoogle(input *model.APILoginForGoogleInput) (output model.APILoginForGoogleOutput)
	APILoginForLine(input *model.APILoginForLineInput) (output model.APILoginForLineOutput)
	APILogout(input *model.APILogoutInput) (output model.APILogoutOutput)
	APICreateRegisterOTP(input *model.APICreateOTPInput) (output model.APICreateRegisterOTPOutput)
	APIRegisterNicknameValidate(input *model.APIRegisterNicknameValidateInput) (output model.APIRegisterNicknameValidateOutput)
	APIRegisterEmailValidate(input *model.APIRegisterEmailValidateInput) (output model.APIRegisterEmailValidateOutput)
	APIRegisterEmailAccountValidate(input *model.APIRegisterEmailAccountValidateInput) (output model.APIRegisterEmailAccountValidateOutput)
	APIRegisterFacebookAccountValidate(input *model.APIRegisterFacebookAccountValidateInput) (output model.APIRegisterFacebookAccountValidateOutput)
	APIRegisterLineAccountValidate(input *model.APIRegisterLineAccountValidateInput) (output model.APIRegisterLineAccountValidateOutput)
	APIRegisterGoogleAccountValidate(input *model.APIRegisterGoogleAccountValidateInput) (output model.APIRegisterGoogleAccountValidateOutput)
}
