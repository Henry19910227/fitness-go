package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/logindto"
	"github.com/Henry19910227/fitness-go/internal/dto/registerdto"
	"github.com/gin-gonic/gin"
)

type Migrate interface {
	Version() (uint, bool, errcode.Error)
	Up() (uint, bool, errcode.Error)
	UpStep(step int) (uint, bool, errcode.Error)
	Down() errcode.Error
	DownStep(step int) errcode.Error
	Force(version int) (uint, bool, errcode.Error)
	Migrate(version uint) (uint, bool, errcode.Error)
}

type Swagger interface {
	WrapHandler() gin.HandlerFunc
}

type Login interface {
	UserLoginByEmail(c *gin.Context, email string, password string) (*logindto.User, string, errcode.Error)
	AdminLoginByEmail(c *gin.Context, email string, password string) (*logindto.Admin, string, errcode.Error)
	UserLogoutByToken(c *gin.Context, token string) errcode.Error
	AdminLogoutByToken(c *gin.Context, token string) errcode.Error
}

type Register interface {
	SendEmailOTP(c *gin.Context, email string) (*registerdto.OTP, errcode.Error)
	EmailRegister(c *gin.Context, otp string, email string, nickname string, password string) (*registerdto.Register, errcode.Error)
	ValidateNicknameDup(c *gin.Context, nickname string) errcode.Error
	ValidateEmailDup(c *gin.Context, email string) errcode.Error
}
