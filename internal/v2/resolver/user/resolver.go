package user

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	userService "github.com/Henry19910227/fitness-go/internal/v2/service/user"
)

type resolver struct {
	userService userService.Service
	otpTool     tool.OTP
}

func New(userService userService.Service, otpTool tool.OTP) Resolver {
	return &resolver{userService: userService, otpTool: otpTool}
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
	userID, err := r.userService.Create(&table)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data.ID = util.PointerInt64(userID)
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
