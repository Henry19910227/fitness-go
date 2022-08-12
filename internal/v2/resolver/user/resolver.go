package user

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/apple_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/google_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iap"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/line_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	receiptModel "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	"github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"strconv"
	"time"
)

type resolver struct {
	userService          user.Service
	receiptService       receipt.Service
	subscribeInfoService user_subscribe_info.Service
	otpTool              otp.Tool
	cryptoTool           crypto.Tool
	redisTool            redis.Tool
	jwtTool              jwt.Tool
	fbLoginTool          fb_login.Tool
	googleLoginTool      google_login.Tool
	appleLoginTool       apple_login.Tool
	lineLoginTool        line_login.Tool
	uploadTool           uploader.Tool
	iapTool              iap.Tool
}

func New(userService user.Service, receiptService receipt.Service,
	subscribeInfoService user_subscribe_info.Service, otpTool otp.Tool,
	cryptoTool crypto.Tool, redisTool redis.Tool,
	jwtTool jwt.Tool, fbLoginTool fb_login.Tool,
	googleLoginTool google_login.Tool, appleLoginTool apple_login.Tool,
	lineLoginTool line_login.Tool, uploadTool uploader.Tool, iapTool iap.Tool) Resolver {
	return &resolver{userService: userService, receiptService: receiptService,
		subscribeInfoService: subscribeInfoService, otpTool: otpTool,
		cryptoTool: cryptoTool, redisTool: redisTool,
		jwtTool: jwtTool, fbLoginTool: fbLoginTool,
		googleLoginTool: googleLoginTool, appleLoginTool: appleLoginTool,
		lineLoginTool: lineLoginTool, uploadTool: uploadTool, iapTool: iapTool}
}

func (r *resolver) APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput) {
	//檢測舊密碼
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.ID)
	findInput.IsDeleted = util.PointerInt(0)
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

func (r *resolver) APIUpdateUserProfile(input *model.APIUpdateUserProfileInput) (output model.APIUpdateUserProfileOutput) {
	//檢查暱稱是否重複
	ok, err := r.nicknameValidate(util.OnNilJustReturnString(input.Body.Nickname, ""))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.BadRequest, errors.New("該暱稱重複").Error())
		return output
	}
	//parser input
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	if err := util.Parser(input.Body, &table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if err := r.userService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	return output
}

func (r *resolver) APIUpdateUserAvatar(input *model.APIUpdateUserAvatarInput) (output model.APIUpdateUserAvatarOutput) {
	fileNamed, err := r.uploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Avatar = util.PointerString(fileNamed)
	if err := r.userService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}

func (r *resolver) APIGetUserProfile(input *model.APIGetUserProfileInput) (output model.APIGetUserProfileOutput) {
	findInput := model.FindInput{}
	findInput.IsDeleted = util.PointerInt(0)
	if err := util.Parser(input, &findInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data, err := r.userService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	outputData := model.APIGetUserProfileData{}
	if err := util.Parser(data, &outputData); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &outputData
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
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	//以access token 取得 fb uid
	fbUid, err := r.fbLoginTool.GetUserID(input.Body.AccessToken)
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

func (r *resolver) APIRegisterForGoogle(input *model.APIRegisterForGoogleInput) (output model.APIRegisterForGoogleOutput) {
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	//以access token 取得 google uid
	guid, err := r.googleLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(guid))
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
	table.AccountType = util.PointerInt(model.Google)
	table.Account = util.PointerString(r.cryptoTool.MD5Encode(guid))
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

func (r *resolver) APIRegisterForApple(input *model.APIRegisterForAppleInput) (output model.APIRegisterForAppleOutput) {
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	//生成 client secret
	secret, err := r.appleLoginTool.GenerateClientSecret(time.Hour)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//以access token 取得 apple uid
	guid, err := r.appleLoginTool.GetUserID(input.Body.AccessToken, secret)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(guid))
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
	table.AccountType = util.PointerInt(model.Apple)
	table.Account = util.PointerString(r.cryptoTool.MD5Encode(guid))
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

func (r *resolver) APIRegisterForLine(input *model.APIRegisterForLineInput) (output model.APIRegisterForLineOutput) {
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	//以access token 取得 client id
	guid, err := r.lineLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(guid))
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
	table.AccountType = util.PointerInt(model.Line)
	table.Account = util.PointerString(r.cryptoTool.MD5Encode(guid))
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
	//獲取user資訊
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(input.Body.Email)
	listInput.Password = util.PointerString(input.Body.Password)
	listInput.IsDeleted = util.PointerInt(0)
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("帳號或密碼錯誤").Error())
		return output
	}
	//更新當前訂閱狀態
	if err := r.updateUserSubscribeInfo(util.OnNilJustReturnInt64(userOutputs[0].ID, 0)); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//取得更新後的user
	findInput := model.FindInput{}
	findInput.ID = userOutputs[0].ID
	findInput.IsDeleted = util.PointerInt(0)
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Trainer"},
		{Field: "UserSubscribeInfo"},
	}
	userOutput, err := r.userService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//產生token
	token, err := r.jwtTool.GenerateUserToken(util.OnNilJustReturnInt64(userOutput.ID, 0))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//設置token過期時間
	key := jwt.UserTokenPrefix + "." + strconv.Itoa(int(util.OnNilJustReturnInt64(userOutput.ID, 0)))
	if err := r.redisTool.SetEX(key, token, r.jwtTool.GetExpire()); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//Parser Output
	data := model.APILoginForEmailData{}
	if err := util.Parser(userOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &data
	output.Token = util.PointerString(token)
	return output
}

func (r *resolver) APILoginForFacebook(input *model.APILoginForFacebookInput) (output model.APILoginForFacebookOutput) {
	//以access token 取得 fb uid
	fbUid, err := r.fbLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(fbUid))
	listInput.IsDeleted = util.PointerInt(0)
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
	output.Token = util.PointerString(token)
	return output
}

func (r *resolver) APILoginForGoogle(input *model.APILoginForGoogleInput) (output model.APILoginForGoogleOutput) {
	//以 access token 取得 google uid
	fbUid, err := r.googleLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(fbUid))
	listInput.IsDeleted = util.PointerInt(0)
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
	data := model.APILoginForGoogleData{}
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
	output.Token = util.PointerString(token)
	return output
}

func (r *resolver) APILoginForLine(input *model.APILoginForLineInput) (output model.APILoginForLineOutput) {
	//以 access token 取得 uid
	fbUid, err := r.lineLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(fbUid))
	listInput.IsDeleted = util.PointerInt(0)
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
	data := model.APILoginForLineData{}
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
	output.Token = util.PointerString(token)
	return output
}

func (r *resolver) APILoginForApple(input *model.APILoginForAppleInput) (output model.APILoginForAppleOutput) {
	//生成 client secret
	secret, err := r.appleLoginTool.GenerateClientSecret(time.Hour)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//以 access token 取得 uid
	uid, err := r.appleLoginTool.GetUserID(input.Body.AccessToken, secret)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//parser input
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
	listInput.IsDeleted = util.PointerInt(0)
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
	data := model.APILoginForAppleData{}
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
	output.Token = util.PointerString(token)
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
	fbUid, err := r.fbLoginTool.GetUserID(input.Body.AccessToken)
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

func (r *resolver) APIRegisterLineAccountValidate(input *model.APIRegisterLineAccountValidateInput) (output model.APIRegisterLineAccountValidateOutput) {
	//以access token 取得 uid
	uid, err := r.lineLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(uid))
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if !ok {
		output.Set(code.DataAlreadyExists, errors.New("該帳號已註冊").Error())
		return output
	}
	output.Set(code.Success, "該帳號可使用")
	return output
}

func (r *resolver) APIRegisterGoogleAccountValidate(input *model.APIRegisterGoogleAccountValidateInput) (output model.APIRegisterGoogleAccountValidateOutput) {
	//以access token 取得 uid
	uid, err := r.googleLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(uid))
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

func (r *resolver) APIRegisterAppleAccountValidate(input *model.APIRegisterAppleAccountValidateInput) (output model.APIRegisterAppleAccountValidateOutput) {
	//檢查帳號是否重複
	ok, err := r.accountValidate(r.cryptoTool.MD5Encode(input.Body.UserIDToken))
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
	listInput.IsDeleted = util.PointerInt(0)
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
	listInput.IsDeleted = util.PointerInt(0)
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
	listInput.IsDeleted = util.PointerInt(0)
	outputs, _, err := r.userService.List(&listInput)
	if err != nil {
		return false, err
	}
	return !(len(outputs) > 0), nil
}

func (r *resolver) updateUserSubscribeInfo(userID int64) error {
	//查詢該用戶訂閱資訊
	subscribeInfoList := subscribeInfoModel.ListInput{}
	subscribeInfoList.UserID = util.PointerInt64(userID)
	subscribeInfoList.Page = 1
	subscribeInfoList.Size = 1
	subscribeInfoList.OrderType = orderBy.DESC
	subscribeInfoList.OrderField = "update_at"
	infoOutputs, _, err := r.subscribeInfoService.List(&subscribeInfoList)
	if err != nil {
		return err
	}
	if len(infoOutputs) == 0 {
		return nil
	}
	//驗證訂閱日期是否過期
	expiresTime, err := time.Parse("2006-01-02 15:04:05", util.OnNilJustReturnString(infoOutputs[0].ExpiresDate, ""))
	if err == nil {
		if time.Now().Unix() > expiresTime.Unix() {
			subscribeInfoTable := subscribeInfoModel.Table{}
			subscribeInfoTable.UserID = util.PointerInt64(userID)
			subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.NoneSubscribe)
			if err := r.subscribeInfoService.Update(&subscribeInfoTable); err != nil {
				return err
			}
		}
	}
	//更新訂閱資訊
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.OrderID = infoOutputs[0].OrderID
	receiptListInput.Page = 1
	receiptListInput.Size = 1
	receiptListInput.OrderType = orderBy.DESC
	receiptListInput.OrderField = "create_at"
	receiptOutputs, _, err := r.receiptService.List(&receiptListInput)
	if err != nil {
		return err
	}
	if len(receiptOutputs) == 0 {
		return nil
	}
	// 驗證iap訂閱狀態
	if util.OnNilJustReturnInt(receiptOutputs[0].PaymentType, 0) == receiptModel.IAP {
		_ = r.updateUserSubscribeInfoForIAP(userID, util.OnNilJustReturnString(receiptOutputs[0].OriginalTransactionID, ""))
	}
	// TODO:驗證iab訂閱狀態
	return nil
}

func (r *resolver) updateUserSubscribeInfoForIAP(userID int64, originalTransactionID string) error {
	token, err := r.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return err
	}
	response, _ := r.iapTool.GetSubscribeAPI(originalTransactionID, token)
	if response != nil {
		if len(response.Data) > 0 {
			if len(response.Data[0].LastTransactions) > 0 {
				subscribeStatus := subscribeInfoModel.NoneSubscribe
				status := response.Data[0].LastTransactions[0].Status
				if status == 1 || status == 3 || status == 4 || status == 5 { // 當前訂閱尚未過期
					subscribeStatus = subscribeInfoModel.ValidSubscribe
				}
				//更新 user_subscribe_info
				subscribeInfoTable := subscribeInfoModel.Table{}
				subscribeInfoTable.UserID = util.PointerInt64(userID)
				subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
				subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(response.Data[0].LastTransactions[0].SignedTransactionInfo.PurchaseDate / 1000).Format("2006-01-02 15:04:05"))
				subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(response.Data[0].LastTransactions[0].SignedTransactionInfo.ExpiresDate / 1000).Format("2006-01-02 15:04:05"))
				if err := r.subscribeInfoService.Update(&subscribeInfoTable); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return nil
}
