package user

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"strconv"
)

type resolver struct {
	userService userService.Service
	otpTool     otp.Tool
	cryptoTool  crypto.Tool
	redisTool   redis.Tool
	jwtTool     jwt.Tool
	fbLoginTool fb_login.Tool
}

func New(userService userService.Service, otpTool otp.Tool, cryptoTool crypto.Tool, redisTool redis.Tool, jwtTool jwt.Tool, fbLoginTool fb_login.Tool) Resolver {
	return &resolver{userService: userService, otpTool: otpTool, cryptoTool: cryptoTool, redisTool: redisTool, jwtTool: jwtTool, fbLoginTool: fbLoginTool}
}

func (r *resolver) APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput) {
	//檢測舊密碼
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.ID)
	data, err := r.userService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if util.OnNilJustReturnString(data.Password, "") != input.Body.OldPassword {
		output.Set(code.PermissionDenied, errors.New("與舊密碼不一致").Error())
		return output
	}
	//修改密碼
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Password = util.PointerString(input.Body.Password)
	if err := r.userService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterForEmail(input *model.APIRegisterForEmailInput) (output model.APIRegisterForEmailOutput) {
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.BadRequest, errors.New("該帳號重複").Error())
		return output
	}
	//檢查暱稱是否重複
	ok, err = r.nicknameValidate(input.Body.Nickname)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.BadRequest, errors.New("該暱稱重複").Error())
		return output
	}
	//檢查Email是否重複
	ok, err = r.emailValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.BadRequest, errors.New("該信箱重複").Error())
		return output
	}
	//創建用戶
	table := model.Table{}
	table.AccountType = util.PointerInt(model.Email)
	table.Account = util.PointerString(input.Body.Email)
	table.Nickname = util.PointerString(input.Body.Nickname)
	table.Email = util.PointerString(input.Body.Email)
	table.Password = util.PointerString(input.Body.Password)
	_, err = r.userService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterForFacebook(input *model.APIRegisterForFacebookInput) (output model.APIRegisterForFacebookOutput) {
	//以access token 取得 fb uid
	fbUid, err := r.fbLoginTool.GetFbUidByAccessToken(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(fbUid))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該帳號重複").Error())
		return output
	}
	//檢查暱稱是否重複
	ok, err = r.nicknameValidate(input.Body.Nickname)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該暱稱重複").Error())
		return output
	}
	//檢查Email是否重複
	ok, err = r.emailValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該信箱重複").Error())
		return output
	}
	//創建用戶
	table := model.Table{}
	table.AccountType = util.PointerInt(model.Facebook)
	table.Account = util.PointerString(r.cryptoTool.MD5Encode(fbUid))
	table.Nickname = util.PointerString(input.Body.Nickname)
	table.Email = util.PointerString(input.Body.Email)
	table.Password = util.PointerString("")
	_, err = r.userService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APILoginForEmail(input *model.APILoginForEmailInput) (output model.APILoginForEmailOutput) {
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(input.Body.Email)
	listInput.Password = util.PointerString(input.Body.Password)
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "UserSubscribeInfo"},
	}
	if err := util.Parser(input.Body, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	datas, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(datas) == 0 {
		output.Set(code.BadRequest, errors.New("帳號或密碼錯誤").Error())
		return output
	}
	data := model.APILoginForEmailData{}
	if err := util.Parser(datas[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//產生token
	token, err := r.jwtTool.GenerateUserToken(util.OnNilJustReturnInt64(data.ID, 0))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//設置token過期時間
	key := jwt.UserTokenPrefix + "." + strconv.Itoa(int(util.OnNilJustReturnInt64(data.ID, 0)))
	if err := r.redisTool.SetEX(key, token, r.jwtTool.GetExpire()); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &data
	output.Token = token
	return output
}

func (r *resolver) APILoginForFacebook(input *model.APILoginForFacebookInput) (output model.APILoginForFacebookOutput) {
	//以access token 取得 fb uid
	fbUid, err := r.fbLoginTool.GetFbUidByAccessToken(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(fbUid))
	listInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "UserSubscribeInfo"},
	}
	if err := util.Parser(input.Body, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	datas, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(datas) == 0 {
		output.Set(code.BadRequest, errors.New("帳號或密碼錯誤").Error())
		return output
	}
	data := model.APILoginForFacebookData{}
	if err := util.Parser(datas[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//產生token
	token, err := r.jwtTool.GenerateUserToken(util.OnNilJustReturnInt64(data.ID, 0))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//設置token過期時間
	key := jwt.UserTokenPrefix + "." + strconv.Itoa(int(util.OnNilJustReturnInt64(data.ID, 0)))
	if err := r.redisTool.SetEX(key, token, r.jwtTool.GetExpire()); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &data
	output.Token = token
	return output
}

func (r *resolver) APILogout(input *model.APILogoutInput) (output model.APILogoutOutput) {
	if err := r.redisTool.Del(jwt.UserTokenPrefix + "." + strconv.Itoa(int(input.ID))); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APICreateRegisterOTP(input *model.APICreateOTPInput) (output model.APICreateRegisterOTPOutput) {
	//產生otp碼
	otp, err := r.otpTool.Generate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	data := model.APICreateRegisterOTPData{}
	data.Code = otp
	output.Data = &data
	return output
}

func (r *resolver) APIRegisterNicknameValidate(input *model.APIRegisterNicknameValidateInput) (output model.APIRegisterNicknameValidateOutput) {
	ok, err := r.nicknameValidate(input.Body.Nickname)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該暱稱不可使用").Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterEmailValidate(input *model.APIRegisterEmailValidateInput) (output model.APIRegisterEmailValidateOutput) {
	ok, err := r.emailValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該信箱不可使用").Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterEmailAccountValidate(input *model.APIRegisterEmailAccountValidateInput) (output model.APIRegisterEmailAccountValidateOutput) {
	ok, err := r.accountValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該帳號不可使用").Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterFacebookAccountValidate(input *model.APIRegisterFacebookAccountValidateInput) (output model.APIRegisterFacebookAccountValidateOutput) {
	//以access token 取得 fb uid
	fbUid, err := r.fbLoginTool.GetFbUidByAccessToken(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(fbUid))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該帳號已註冊").Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) nicknameValidate(nickname string) (bool, error) {
	//檢查帳號是否重複
	listInput := model.ListInput{}
	listInput.Nickname = util.PointerString(nickname)
	outputs, _, err := r.userService.List(&listInput)
	if err != nil {
		return false, err
	}
	return !(len(outputs) > 0), nil
}

func (r *resolver) emailValidate(email string) (bool, error) {
	//檢查帳號是否重複
	listInput := model.ListInput{}
	listInput.Email = util.PointerString(email)
	outputs, _, err := r.userService.List(&listInput)
	if err != nil {
		return false, err
	}
	return !(len(outputs) > 0), nil
}

func (r *resolver) accountValidate(account string) (bool, error) {
	//檢查帳號是否重複
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(account)
	outputs, _, err := r.userService.List(&listInput)
	if err != nil {
		return false, err
	}
	return !(len(outputs) > 0), nil
}
