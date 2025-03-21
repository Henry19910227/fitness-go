package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/dto/registerdto"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
)

type register struct {
	Base
	userRepo repository.User
	logger    handler.Logger
	jwtTool   tool.JWT
	otpTool   tool.OTP
	viperTool  *viper.Viper
	errHandler errcode.Handler
}

func NewRegister(
	userRepo repository.User,
	logger handler.Logger,
	jwtTool tool.JWT,
	otpTool tool.OTP,
	viperTool  *viper.Viper,
	errHandler errcode.Handler) Register {

	return &register{
		userRepo: userRepo,
		logger: logger,
		jwtTool: jwtTool,
		otpTool: otpTool,
		viperTool: viperTool,
		errHandler: errHandler}
}

func (r *register) SendEmailOTP(c *gin.Context, email string) (*registerdto.OTP, errcode.Error) {
	//產生otp碼
	otp, err := r.otpTool.Generate(email)
	if err != nil {
		r.logger.Set(c, handler.Error, "otp tool", r.errHandler.SendOTPFailure().Code(), err.Error())
		return nil, r.errHandler.SendOTPFailure()
	}
	//如果是debug模式就不發送簡訊，並回傳otp
	if r.viperTool.GetString("Server.RunMode") == "debug" {
		return &registerdto.OTP{Code: otp}, nil
	}
	//暫時只回傳otp不寄信
	return &registerdto.OTP{Code: otp}, nil
}

func (r *register) EmailRegister(c *gin.Context, otp string, email string, nickname string, password string) (*registerdto.Register, errcode.Error) {
	//驗證OTP
	if !r.otpTool.Validate(otp, email) {
		return nil, r.errHandler.OTPInvalid()
	}
	//創建用戶
	uid, err := r.userRepo.CreateUser(1, email, nickname, password)
	if err != nil {
		//有重複的欄位資料
		if r.MysqlDuplicateEntry(err) {
			if strings.Contains(err.Error(), "account") {
				return nil, r.errHandler.Custom(9004,errors.New("重複的帳戶名稱"))
			}
			if strings.Contains(err.Error(), "nickname") {
				return nil, r.errHandler.Custom(9004,errors.New("重複的用戶暱稱"))
			}
			return nil, r.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
		r.logger.Set(c, handler.Error, "UserRepo", r.errHandler.SystemError().Code(), err.Error())
		return nil, r.errHandler.SystemError()
	}
	return &registerdto.Register{UserID: uid}, nil
}

func (r *register) ValidateNicknameDup(c *gin.Context, nickname string) errcode.Error {
	_, err := r.userRepo.FindUserIDByNickname(nickname)
	//該帳號已存在
	if err == nil {
		return r.errHandler.NicknameDuplicate()
	}
	//該帳號可使用
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil
	}
	//不明原因錯誤
	r.logger.Set(c, handler.Error, "UserRepo", r.errHandler.SystemError().Code(), err.Error())
	return r.errHandler.SystemError()
}

func (r *register) ValidateEmailDup(c *gin.Context, email string) errcode.Error {
	_, err := r.userRepo.FindUserIDByEmail(email)
	//該信箱已存在
	if err == nil {
		return r.errHandler.EmailDuplicate()
	}
	//該信箱可使用
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil
	}
	//不明原因錯誤
	r.logger.Set(c, handler.Error, "UserRepo", r.errHandler.SystemError().Code(), err.Error())
	return r.errHandler.SystemError()
}
