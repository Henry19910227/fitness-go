package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
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
	regErr    errcode.Register
}

func NewRegister(
	userRepo   repository.User,
	logger     handler.Logger,
	jwtTool    tool.JWT,
	otpTool    tool.OTP,
	viperTool  *viper.Viper,
	regErr     errcode.Register) Register {

	return &register{
		userRepo: userRepo,
		logger: logger,
		jwtTool: jwtTool,
		otpTool: otpTool,
		viperTool: viperTool,
		regErr: regErr}
}

func (r *register) EmailRegister(c *gin.Context, otp string, email string, nickname string, password string) errcode.Error {
	panic("implement me")
}
