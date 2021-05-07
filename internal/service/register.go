package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/registerdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type register struct {
	Base
	userRepo  repository.User
	logger    handler.Logger
	jwtTool   tool.JWT
	otpTool   tool.OTP
	viperTool *viper.Viper
	errHandler errcode.Handler
}

func NewRegister(
	userRepo   repository.User,
	logger     handler.Logger,
	jwtTool    tool.JWT,
	otpTool    tool.OTP,
	viperTool  *viper.Viper,
	errHandler  errcode.Handler) Register {

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

func (r *register) EmailRegister(c *gin.Context, otp string, email string, nickname string, password string) (*registerdto.RegisterResult, errcode.Error) {
	//驗證OTP
	if !r.otpTool.Validate(otp, email) {
		return nil, r.errHandler.OTPInvalid()
	}
	//創建用戶
	uid, err := r.userRepo.CreateUser(1, email, nickname, password)
	if err != nil {
		//有重複的欄位資料
		if r.MysqlDuplicateEntry(err) {
			return nil, r.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
		r.logger.Set(c, handler.Error, "UserRepo", r.errHandler.SystemError().Code(), err.Error())
		return nil, r.errHandler.SystemError()
	}
	return &registerdto.RegisterResult{UserID: uid}, nil
}
