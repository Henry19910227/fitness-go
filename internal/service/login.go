package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/logindto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type login struct {
	Base
	adminRepo repository.Admin
	userRepo  repository.User
	ssoHandler handler.SSO
	jwtTool   tool.JWT
	logger handler.Logger
	errHandler errcode.Handler
}

func NewLogin(adminRepo repository.Admin,
	userRepo  repository.User,
	ssoHandler handler.SSO,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Login {
	return &login{adminRepo: adminRepo,
		userRepo: userRepo,
		ssoHandler: ssoHandler,
		logger: logger,
		jwtTool: jwtTool,
		errHandler: errHandler}
}

func (l *login) UserLoginByEmail(c *gin.Context, email string, password string) (*logindto.User, string, errcode.Error) {
	//從db查詢用戶
	var user logindto.User
	if err := l.userRepo.FindUserByAccountAndPassword(email, password, &user); err != nil {
		//查無此人
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", l.errHandler.LoginFailure()
		}
		//不明原因錯誤
		l.logger.Set(c, handler.Error, "UserRepo", l.errHandler.SystemError().Code(), err.Error())
		return nil, "", l.errHandler.SystemError()
	}
	if user.Birthday == "0000-01-01" {
		user.Birthday = ""
	}
	//生成 user token
	token, err := l.ssoHandler.GenerateUserToken(user.ID)
	if err != nil {
		l.logger.Set(c, handler.Error, "SsoHandler", l.errHandler.SystemError().Code(), err.Error())
		return nil, "", l.errHandler.SystemError()
	}
	//設置上線狀態
	if err := l.ssoHandler.RenewOnlineStatus(token); err != nil {
		l.logger.Set(c, handler.Error, "SsoHandler", l.errHandler.SystemError().Code(), err.Error())
		return nil, "", l.errHandler.SystemError()
	}
	return &user, token, nil
}

func (l *login) AdminLoginByEmail(c *gin.Context, email string, password string) (*logindto.Admin, string, errcode.Error) {
	uid, err := l.adminRepo.GetAdminID(email, password)
	if err != nil {
		//查無此人
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", l.errHandler.LoginFailure()
		}
		//不明原因錯誤
		l.logger.Set(c, handler.Error, "AdminRepo", l.errHandler.SystemError().Code(), err.Error())
		return nil, "", l.errHandler.SystemError()
	}
	var admin logindto.Admin
	if err := l.adminRepo.GetAdmin(uid, &admin); err != nil {
		l.logger.Set(c, handler.Error, "AdminRepo", l.errHandler.SystemError().Code(), err.Error())
		return nil, "", l.errHandler.SystemError()
	}
	//生成 Admin Token
	token, e := l.ssoHandler.GenerateAdminToken(uid, admin.Lv)
	if e != nil {
		l.logger.Set(c, handler.Error, "SSOHandler", l.errHandler.SystemError().Code(), e.Error())
		return nil, "", l.errHandler.SystemError()
	}
	return &admin, token, nil
}

func (l *login) Logout(c *gin.Context, token string) errcode.Error {
	if err := l.ssoHandler.SetOfflineStatus(token); err != nil {
		l.logger.Set(c, handler.Error, "SSOHandler", l.errHandler.SystemError().Code(), err.Error())
		return l.errHandler.SystemError()
	}
	if err := l.ssoHandler.ResignUserToken(token); err != nil {
		l.logger.Set(c, handler.Error, "SSOHandler", l.errHandler.SystemError().Code(), err.Error())
		return l.errHandler.SystemError()
	}
	return nil
}


func (l *login) AdminLogoutByToken(c *gin.Context, token string) errcode.Error {
	if err := l.ssoHandler.ResignAdminToken(token); err != nil {
		l.logger.Set(c, handler.Error, "SSOHandler", l.errHandler.SystemError().Code(), err.Error())
		return l.errHandler.SystemError()
	}
	return nil
}
