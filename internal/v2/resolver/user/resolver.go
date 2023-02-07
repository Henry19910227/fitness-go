package user

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/build"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/apple_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/crypto"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/google_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iab"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/iap"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/jwt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/line_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/mail"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/otp"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	receiptModel "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_course_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_user"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_get_cms_users"
	"github.com/Henry19910227/fitness-go/internal/v2/model/user/api_update_cms_user"
	subscribeInfoModel "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/receipt"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user"
	"github.com/Henry19910227/fitness-go/internal/v2/service/user_subscribe_info"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type resolver struct {
	userService          user.Service
	receiptService       receipt.Service
	subscribeInfoService user_subscribe_info.Service
	courseService        course.Service
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
	iabTool              iab.Tool
	mailTool             mail.Tool
}

func New(userService user.Service, receiptService receipt.Service,
	subscribeInfoService user_subscribe_info.Service, courseService course.Service,
	otpTool otp.Tool, cryptoTool crypto.Tool, redisTool redis.Tool,
	jwtTool jwt.Tool, fbLoginTool fb_login.Tool,
	googleLoginTool google_login.Tool, appleLoginTool apple_login.Tool,
	lineLoginTool line_login.Tool, uploadTool uploader.Tool,
	iapTool iap.Tool, iabTool iab.Tool, mailTool mail.Tool) Resolver {
	return &resolver{userService: userService, receiptService: receiptService,
		subscribeInfoService: subscribeInfoService, courseService: courseService, otpTool: otpTool,
		cryptoTool: cryptoTool, redisTool: redisTool,
		jwtTool: jwtTool, fbLoginTool: fbLoginTool,
		googleLoginTool: googleLoginTool, appleLoginTool: appleLoginTool,
		lineLoginTool: lineLoginTool, uploadTool: uploadTool,
		iapTool: iapTool, iabTool: iabTool, mailTool: mailTool}
}

func (r *resolver) APIGetCMSCourseUsers(input *api_get_cms_course_users.Input) (output api_get_cms_course_users.Output) {
	// 查詢課表使用者
	userListInput := model.ListInput{}
	userListInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN user_course_assets ON users.id = user_course_assets.user_id"},
	}
	userListInput.Wheres = []*whereModel.Where{
		{Query: "user_course_assets.course_id = ?", Args: []interface{}{input.Uri.CourseID}},
	}
	userListInput.Preloads = []*preloadModel.Preload{
		{Field: "UserCourseAsset", Conditions: []interface{}{"course_id = ?", input.Uri.CourseID}},
	}
	userListInput.Orders = []*orderBy.Order{
		{Value: fmt.Sprintf("user_course_assets.%s %s", input.Query.OrderField, input.Query.OrderType)},
	}
	userListInput.Size = input.Query.Size
	userListInput.Page = input.Query.Page
	userOutputs, page, err := r.userService.List(&userListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parse output
	data := api_get_cms_course_users.Data{}
	if err := util.Parser(userOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetCMSUser(input *api_get_cms_user.Input) (output api_get_cms_user.Output) {
	// 查詢課表使用者
	findUserInput := model.FindInput{}
	findUserInput.ID = util.PointerInt64(input.Uri.UserID)
	userOutput, err := r.userService.Find(&findUserInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parse output
	data := api_get_cms_user.Data{}
	if err := util.Parser(userOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetCMSUsers(input *api_get_cms_users.Input) (output api_get_cms_users.Output) {
	if input.Query.OrderField == nil {
		input.Query.OrderField = util.PointerString("create_at")
	}
	if input.Query.OrderType == nil {
		input.Query.OrderType = util.PointerString(orderBy.DESC)
	}
	wheres := make([]*whereModel.Where, 0)
	if input.Query.Nickname != nil {
		wheres = append(wheres, &whereModel.Where{Query: "users.nickname LIKE ?", Args: []interface{}{"%" + *input.Query.Nickname + "%"}})
	}
	if input.Query.Email != nil {
		wheres = append(wheres, &whereModel.Where{Query: "users.email = ?", Args: []interface{}{"%" + *input.Query.Email + "%"}})
	}
	// 查詢用戶
	userListInput := model.ListInput{}
	userListInput.ID = input.Query.UserID
	userListInput.UserStatus = input.Query.UserStatus
	userListInput.UserType = input.Query.UserType
	userListInput.Wheres = wheres
	userListInput.Size = input.Query.Size
	userListInput.Page = input.Query.Page
	userListInput.OrderField = *input.Query.OrderField
	userListInput.OrderType = *input.Query.OrderType
	userListInput.Size = input.Query.Size
	userListInput.Page = input.Query.Page
	userOutputs, page, err := r.userService.List(&userListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parse output
	data := api_get_cms_users.Data{}
	if err := util.Parser(userOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateCMSUser(input *api_update_cms_user.Input) (output api_update_cms_user.Output) {
	userTable := model.Table{}
	userTable.ID = util.PointerInt64(input.Uri.UserID)
	if input.Body.Password != nil {
		userTable.Password = util.PointerString(r.cryptoTool.MD5Encode(*input.Body.Password))
	}
	userTable.UserStatus = input.Body.UserStatus
	if err := r.userService.Update(&userTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	return output
}

func (r *resolver) APIUpdatePassword(input *model.APIUpdatePasswordInput) (output model.APIUpdatePasswordOutput) {
	//檢測舊密碼
	findInput := model.FindInput{}
	findInput.ID = util.PointerInt64(input.ID)
	findInput.IsDeleted = util.PointerInt(0)
	userOutput, err := r.userService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	oldPassword := r.cryptoTool.MD5Encode(input.Body.OldPassword)
	if util.OnNilJustReturnString(userOutput.Password, "") != oldPassword {
		output.Set(code.PermissionDenied, errors.New("與舊密碼不一致").Error())
		return output
	}
	//修改密碼
	table := model.Table{}
	table.ID = util.PointerInt64(input.ID)
	table.Password = util.PointerString(r.cryptoTool.MD5Encode(input.Body.Password))
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

func (r *resolver) APIGetAppleRefreshToken(input *model.APIGetAppleRefreshTokenInput) (output model.APIGetAppleRefreshTokenOutput) {
	//生成 client secret
	secret, err := r.appleLoginTool.GenerateClientSecret(time.Hour)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取 refresh token
	refreshToken, err := r.appleLoginTool.APIGetRefreshToken(input.Body.AccessToken, secret)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &model.APIGetAppleRefreshTokenData{
		RefreshToken: refreshToken,
	}
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
	table.Password = util.PointerString(r.cryptoTool.MD5Encode(input.Body.Password))
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
	//以 refresh token 取得 uid
	uid, err := r.appleLoginTool.APIGetUserID(input.Body.RefreshToken, secret)
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
	table.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
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
	listInput.Password = util.PointerString(r.cryptoTool.MD5Encode(input.Body.Password))
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
	uid, err := r.fbLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取user資訊
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
	listInput.IsDeleted = util.PointerInt(0)
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("查無此用戶").Error())
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
	data := model.APILoginForFacebookData{}
	if err := util.Parser(userOutput, &data); err != nil {
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
	uid, err := r.googleLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取user資訊
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
	listInput.IsDeleted = util.PointerInt(0)
	if err := util.Parser(input.Body, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("查無此用戶").Error())
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
	data := model.APILoginForGoogleData{}
	if err := util.Parser(userOutput, &data); err != nil {
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
	uid, err := r.lineLoginTool.GetUserID(input.Body.AccessToken)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取user資訊
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
	listInput.IsDeleted = util.PointerInt(0)
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("查無此用戶").Error())
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
	data := model.APILoginForLineData{}
	if err := util.Parser(userOutput, &data); err != nil {
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
	//以 refresh token 取得 uid
	uid, err := r.appleLoginTool.APIGetUserID(input.Body.RefreshToken, secret)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//獲取user資訊
	listInput := model.ListInput{}
	listInput.Account = util.PointerString(r.cryptoTool.MD5Encode(uid))
	listInput.IsDeleted = util.PointerInt(0)
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, errors.New("查無此用戶").Error())
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
	//Parser output
	data := model.APILoginForAppleData{}
	if err := util.Parser(userOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = &data
	output.Token = util.PointerString(token)
	return output
}

func (r *resolver) APILogout(input *model.APILogoutInput) (output model.APILogoutOutput) {
	// 移除登入狀態
	if err := r.redisTool.Del(jwt.UserTokenPrefix + "." + strconv.Itoa(int(input.ID))); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 移除 device token
	userTable := model.Table{}
	userTable.ID = util.PointerInt64(input.ID)
	userTable.DeviceToken = util.PointerString("")
	if err := r.userService.Update(&userTable); err != nil {
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
	if build.RunMode() == "production" {
		data.Code = ""
		if err := r.mailTool.Send(input.Body.Email, "驗證碼", otp); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
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
	//生成 client secret
	secret, err := r.appleLoginTool.GenerateClientSecret(time.Hour)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	//以 refresh token 取得 uid
	uid, err := r.appleLoginTool.APIGetUserID(input.Body.RefreshToken, secret)
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
	output.Set(code.Success, "該帳號可註冊")
	return output
}

func (r *resolver) APICreateResetOTP(input *model.APICreateResetOTPInput) (output model.APICreateResetOTPOutput) {
	//驗證權限
	listInput := model.ListInput{}
	listInput.Email = util.PointerString(input.Body.Email)
	listInput.Page = util.PointerInt(1)
	listInput.Size = util.PointerInt(1)
	listInput.OrderField = "create_at"
	listInput.OrderType = orderBy.DESC
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, "查無此帳戶")
		return output
	}
	if util.OnNilJustReturnInt(userOutputs[0].AccountType, 0) != model.Email {
		output.Set(code.PermissionDenied, "非信箱註冊帳戶")
		return output
	}
	//產生otp碼
	otp, err := r.otpTool.Generate(input.Body.Email)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APICreateResetOTPData{}
	data.Code = otp
	output.SetStatus(code.Success)
	output.Data = &data
	if build.RunMode() == "production" {
		data.Code = ""
		if err := r.mailTool.Send(input.Body.Email, "驗證碼", otp); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
	}
	return output
}

func (r *resolver) APIResetOTPValidate(input *model.APIResetOTPValidateInput) (output model.APIResetOTPValidateOutput) {
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, errors.New("無效的驗證碼").Error())
		return output
	}
	output.Set(code.Success, "驗證成功!")
	return output
}

func (r *resolver) APIUpdateResetPassword(input *model.APIUpdateResetPasswordInput) (output model.APIUpdateResetPasswordOutput) {
	//驗證權限
	listInput := model.ListInput{}
	listInput.Email = util.PointerString(input.Body.Email)
	listInput.IsDeleted = util.PointerInt(0)
	listInput.Page = util.PointerInt(1)
	listInput.Size = util.PointerInt(1)
	listInput.OrderField = "create_at"
	listInput.OrderType = orderBy.DESC
	userOutputs, _, err := r.userService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, "查無此帳戶")
		return output
	}
	if util.OnNilJustReturnInt(userOutputs[0].AccountType, 0) != model.Email {
		output.Set(code.PermissionDenied, "無法修改非Email註冊的帳戶密碼")
		return output
	}
	//檢查驗證碼
	if !r.otpTool.Validate(input.Body.OTPCode, input.Body.Email) {
		output.Set(code.BadRequest, "無效的驗證碼")
		return output
	}
	//修改密碼
	userTable := model.Table{}
	userTable.ID = userOutputs[0].ID
	userTable.Password = util.PointerString(r.cryptoTool.MD5Encode(input.Body.Password))
	if err := r.userService.Update(&userTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "驗證成功!")
	return output
}

func (r *resolver) APIDeleteUser(ctx *gin.Context, tx *gorm.DB, input *model.APIDeleteUserInput) (output model.APIDeleteUserOutput) {
	defer tx.Rollback()
	// 查詢該用戶
	listInput := model.ListInput{}
	listInput.ID = util.PointerInt64(input.UserID)
	listInput.IsDeleted = util.PointerInt(0)
	listInput.Page = util.PointerInt(1)
	listInput.Size = util.PointerInt(1)
	listInput.OrderField = "create_at"
	listInput.OrderType = orderBy.DESC
	userOutputs, _, err := r.userService.Tx(tx).List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(userOutputs) == 0 {
		output.Set(code.BadRequest, "查無此用戶")
		return output
	}
	// 將該用戶更新為刪除狀態
	table := model.Table{}
	table.ID = util.PointerInt64(input.UserID)
	table.IsDeleted = util.PointerInt(1)
	if err := r.userService.Tx(tx).Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 查詢該用戶上架的付費課表
	courseListInput := courseModel.ListInput{}
	courseListInput.UserID = util.PointerInt64(input.UserID)
	courseListInput.SaleType = util.PointerInt(courseModel.SaleTypeCharge)
	courseListInput.CourseStatus = util.PointerInt(courseModel.Sale)
	courseOutputs, _, err := r.courseService.Tx(tx).List(&courseListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 將付費課表改為下架狀態
	courseTables := make([]*courseModel.Table, 0)
	for _, courseOutput := range courseOutputs {
		courseTable := courseModel.Table{}
		courseTable.ID = courseOutput.ID
		courseTable.CourseStatus = util.PointerInt(courseModel.Remove)
		courseTables = append(courseTables, &courseTable)
	}
	if err := r.courseService.Tx(tx).Updates(courseTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// 寄信
	if err := r.mailTool.Send(util.OnNilJustReturnString(userOutputs[0].Email, ""), "健身帳戶刪除通知信", "您已刪除健身帳戶"); err != nil {
		logger.Shared().Error(ctx, err.Error())
	}
	// 將token失效
	if err := r.redisTool.Del(jwt.UserTokenPrefix + "." + strconv.Itoa(int(input.UserID))); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "刪除成功!")
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

func (r *resolver) updateUserSubscribeInfo(userID int64) error {
	//查詢該用戶訂閱資訊
	subscribeInfoList := subscribeInfoModel.ListInput{}
	subscribeInfoList.UserID = util.PointerInt64(userID)
	subscribeInfoList.Page = util.PointerInt(1)
	subscribeInfoList.Size = util.PointerInt(1)
	subscribeInfoList.OrderType = orderBy.DESC
	subscribeInfoList.OrderField = "update_at"
	infoOutputs, _, err := r.subscribeInfoService.List(&subscribeInfoList)
	if err != nil {
		return err
	}
	if len(infoOutputs) == 0 {
		return nil
	}
	// 驗證訂閱狀態是否正常(帳號在效期內且狀態為訂閱狀態 or 帳號過期且狀態為非訂閱狀態)
	nowUnix := time.Now().Unix()
	expiresUnix := util.DateStringToTime(util.OnNilJustReturnString(infoOutputs[0].ExpiresDate, "")).Unix()
	if nowUnix <= expiresUnix && util.OnNilJustReturnInt(infoOutputs[0].Status, 0) > 0 {
		return nil
	}
	if nowUnix >= expiresUnix && util.OnNilJustReturnInt(infoOutputs[0].Status, 0) == 0 {
		return nil
	}
	// 帳號訂閱狀態不正常則同步線上 iap or iab 訂閱狀態
	subscribeInfoTable := subscribeInfoModel.Table{}
	subscribeInfoTable.UserID = util.PointerInt64(userID)
	subscribeInfoTable.Status = util.PointerInt(subscribeInfoModel.NoneSubscribe)
	if err := r.subscribeInfoService.Update(&subscribeInfoTable); err != nil {
		return err
	}
	// 同步iap訂閱狀態
	if util.OnNilJustReturnInt(infoOutputs[0].PaymentType, 0) == receiptModel.IAP {
		_ = r.updateUserSubscribeInfoForIAP(infoOutputs[0])
	}
	// 同步iab訂閱狀態
	if util.OnNilJustReturnInt(infoOutputs[0].PaymentType, 0) == receiptModel.IAB {
		_ = r.updateUserSubscribeInfoForIAB(infoOutputs[0])
	}
	return nil
}

func (r *resolver) updateUserSubscribeInfoForIAP(subscribeInfo *subscribeInfoModel.Output) error {
	//查詢收據資料
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.OrderID = subscribeInfo.OrderID
	receiptListInput.PaymentType = util.PointerInt(receiptModel.IAP)
	receiptListInput.Page = util.PointerInt(1)
	receiptListInput.Size = util.PointerInt(1)
	receiptListInput.OrderType = orderBy.DESC
	receiptListInput.OrderField = "create_at"
	receiptOutputs, _, err := r.receiptService.List(&receiptListInput)
	if err != nil {
		return err
	}
	if len(receiptOutputs) == 0 {
		return nil
	}
	// 查詢並同步當前線上IAP訂閱狀態
	token, err := r.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return err
	}
	response, _ := r.iapTool.GetSubscribeAPI(util.OnNilJustReturnString(receiptOutputs[0].OriginalTransactionID, ""), token)
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
				subscribeInfoTable.UserID = subscribeInfo.UserID
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

func (r *resolver) updateUserSubscribeInfoForIAB(subscribeInfo *subscribeInfoModel.Output) error {
	//查詢收據資料
	receiptListInput := receiptModel.ListInput{}
	receiptListInput.Wheres = []*whereModel.Where{
		{Query: "LENGTH(receipts.receipt_token) > 0"},
	}
	receiptListInput.OrderID = subscribeInfo.OrderID
	receiptListInput.PaymentType = util.PointerInt(receiptModel.IAB)
	receiptListInput.Page = util.PointerInt(1)
	receiptListInput.Size = util.PointerInt(1)
	receiptListInput.OrderType = orderBy.DESC
	receiptListInput.OrderField = "create_at"
	receiptOutputs, _, err := r.receiptService.List(&receiptListInput)
	if err != nil {
		return err
	}
	if len(receiptOutputs) == 0 {
		return nil
	}
	//產出 auth token
	oauthToken, err := r.iabTool.GenerateGoogleOAuth2Token(time.Hour)
	if err != nil {
		return err
	}
	//獲取API Token
	token, err := r.iabTool.APIGetGooglePlayToken(oauthToken)
	if err != nil {
		return err
	}
	// 查詢並同步當前線上IAB訂閱狀態
	productID := util.OnNilJustReturnString(receiptOutputs[0].ProductID, "")
	receiptToken := util.OnNilJustReturnString(receiptOutputs[0].ReceiptToken, "")
	response, err := r.iabTool.APIGetSubscription(productID, receiptToken, token)
	if err != nil {
		return err
	}
	if response != nil {
		subscribeStatus := subscribeInfoModel.NoneSubscribe
		if response.PaymentState == 1 || response.PaymentState == 2 { // 當前訂閱尚未過期
			subscribeStatus = subscribeInfoModel.ValidSubscribe
		}
		//更新 user_subscribe_info
		startTimeMillis, err := strconv.ParseInt(response.StartTimeMillis, 10, 64)
		if err != nil {
			return err
		}
		expiryTimeMillis, err := strconv.ParseInt(response.ExpiryTimeMillis, 10, 64)
		if err != nil {
			return err
		}
		subscribeInfoTable := subscribeInfoModel.Table{}
		subscribeInfoTable.UserID = subscribeInfo.UserID
		subscribeInfoTable.Status = util.PointerInt(subscribeStatus)
		subscribeInfoTable.StartDate = util.PointerString(util.UnixToTime(startTimeMillis / 1000).Format("2006-01-02 15:04:05"))
		subscribeInfoTable.ExpiresDate = util.PointerString(util.UnixToTime(expiryTimeMillis / 1000).Format("2006-01-02 15:04:05"))
		if err := r.subscribeInfoService.Update(&subscribeInfoTable); err != nil {
			return err
		}
	}
	return nil
}
