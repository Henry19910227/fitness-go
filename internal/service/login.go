package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type login struct {
	Base
	adminRepo         repository.Admin
	userRepo          repository.User
	trainerRepo       repository.Trainer
	albumRepo         repository.TrainerAlbum
	cerRepo           repository.Certificate
	subscribeInfoRepo repository.UserSubscribeInfo
	ssoHandler        handler.SSO
	jwtTool           tool.JWT
	logger            handler.Logger
	errHandler        errcode.Handler
}

func NewLogin(adminRepo repository.Admin,
	userRepo repository.User,
	trainerRepo repository.Trainer,
	albumRepo repository.TrainerAlbum,
	cerRepo repository.Certificate,
	subscribeInfoRepo repository.UserSubscribeInfo,
	ssoHandler handler.SSO,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Login {
	return &login{adminRepo: adminRepo,
		userRepo:          userRepo,
		trainerRepo:       trainerRepo,
		albumRepo:         albumRepo,
		cerRepo:           cerRepo,
		subscribeInfoRepo: subscribeInfoRepo,
		ssoHandler:        ssoHandler,
		logger:            logger,
		jwtTool:           jwtTool,
		errHandler:        errHandler}
}

func (l *login) UserLoginByEmail(c *gin.Context, email string, password string) (*dto.User, string, errcode.Error) {
	//從db查詢用戶
	var user dto.User
	if err := l.userRepo.FindUserByAccountAndPassword(email, password, &user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errcode.LoginFailure
		}
		return nil, "", l.errHandler.Set(c, "user repo", err)
	}
	if user.Birthday == "0000-01-01" {
		user.Birthday = ""
	}
	//獲取教練資訊
	data, err := l.trainerRepo.FindTrainer(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", l.errHandler.Set(c, "trainer repo", err)
	}
	if data != nil {
		trainer := dto.NewTrainer(data)
		if err := l.albumRepo.FindAlbumPhotosByUID(user.ID, &trainer.TrainerAlbumPhotos); err != nil {
			return nil, "", l.errHandler.Set(c, "trainer album repo", err)
		}
		if err := l.cerRepo.FindCertificatesByUID(user.ID, &trainer.Certificates); err != nil {
			return nil, "", l.errHandler.Set(c, "cer repo", err)
		}
		user.TrainerInfo = &trainer
	}
	//獲取訂閱資訊
	subscribeInfoData, err := l.subscribeInfoRepo.FindSubscribeInfo(user.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", l.errHandler.Set(c, "user subscribe info repo", err)
	}
	if subscribeInfoData != nil {
		user.SubscribeInfo = &dto.UserSubscribeInfo{
			Status:      subscribeInfoData.Status,
			StartDate:   subscribeInfoData.StartDate,
			ExpiresDate: subscribeInfoData.ExpiresDate,
		}
	}
	//生成 user token
	token, err := l.ssoHandler.GenerateUserToken(user.ID)
	if err != nil {
		return nil, "", l.errHandler.Set(c, "sso handler", err)
	}
	return &user, token, nil
}

func (l *login) AdminLoginByEmail(c *gin.Context, email string, password string) (*dto.Admin, string, errcode.Error) {
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
	var admin dto.Admin
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
	l.logger.Set(c, handler.Error, "Admin Login", 0, "Admin Login Success!")
	return &admin, token, nil
}

func (l *login) UserLogoutByToken(c *gin.Context, token string) errcode.Error {
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
