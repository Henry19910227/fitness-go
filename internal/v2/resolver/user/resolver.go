package user

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
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
}

func New(userService userService.Service, otpTool otp.Tool, cryptoTool crypto.Tool, redisTool redis.Tool, jwtTool jwt.Tool) Resolver {
	return &resolver{userService: userService, otpTool: otpTool, cryptoTool: cryptoTool, redisTool: redisTool, jwtTool: jwtTool}
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

func (r *resolver) APICreateRegisterOTP(input *model.APICreateRegisterOTPInput) (output model.APICreateRegisterOTPOutput) {
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
		output.Set(code.BadRequest, errors.New("該暱稱不可使用").Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIRegisterAccountValidate(input *model.APIRegisterAccountValidateInput) (output model.APIRegisterAccountValidateOutput) {
	ok, err := r.accountValidate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.BadRequest, errors.New("該帳號不可使用").Error())
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
