package service

//import (
//	"errors"
//	"github.com/Henry19910227/fitness-go/errcode"
//	"github.com/Henry19910227/fitness-go/internal/handler"
//	"github.com/Henry19910227/fitness-go/internal/model/admindata"
//	"github.com/Henry19910227/fitness-go/internal/model/logindata"
//	"github.com/Henry19910227/fitness-go/internal/repository"
//	"github.com/Henry19910227/fitness-go/internal/tool"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//)
//
//type login struct {
//	Base
//	userRepo  repository.User
//	albumRepo repository.Album
//	adminRepo repository.Admin
//	ssoHandler handler.SSO
//	jwtTool   tool.JWT
//	loginErr  errcode.Login
//}
//
//func NewLogin(userRepo repository.User,
//	adminRepo repository.Admin,
//	fbHandler handler.FbLogin,
//	ssoHandler handler.SSO,
//	logger handler.Logger,
//	jwtTool tool.JWT,
//	loginErr errcode.Login) Login {
//	return &login{userRepo: userRepo,
//		adminRepo: adminRepo,
//		fbHandler: fbHandler,
//		ssoHandler: ssoHandler,
//		logger: logger,
//		jwtTool: jwtTool,
//		loginErr: loginErr}
//}
//
//func (l *login) FBLogin(c *gin.Context, accessToken string, role int) (*logindata.User, string, errcode.Error) {
//	// 驗證fb access token並且回傳 fb uid
//	fbUID, e := l.fbHandler.VerifyFBToken(accessToken)
//	if e != nil {
//		l.logger.Set(c, handler.Error, "FBLoginHandler", l.loginErr.SystemError().Code(), e.Error())
//		return nil, "", l.loginErr.InvalidThirdParty()
//	}
//	// 取得 User
//	user, token, err := l.login(c, fbUID, "", role)
//	if err != nil {
//		return nil, "", err
//	}
//	return user, token, nil
//}
//
//func (l *login) MobileLogin(c *gin.Context, mobile string, password string, role int) (*logindata.User, string, errcode.Error) {
//	// 調用 Service
//	user, token, err := l.login(c, mobile, password, role)
//	if err != nil {
//		return nil, "", err
//	}
//	return user, token, nil
//}
//
//func (l *login) LoginForAdmin(c *gin.Context, email string, password string) (*admindata.Admin, string, errcode.Error) {
//	uid, err := l.adminRepo.GetAdminID(email, password)
//	if err != nil {
//		//查無此人
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, "", l.loginErr.LoginFailure()
//		}
//		//不明原因錯誤
//		l.logger.Set(c, handler.Error, "AdminRepo", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	var admin admindata.Admin
//	if err := l.adminRepo.GetAdmin(uid, &admin); err != nil {
//		l.logger.Set(c, handler.Error, "AdminRepo", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	//生成 Admin Token
//	token, e := l.ssoHandler.GenerateAdminToken(uid, admin.Lv)
//	if e != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), e.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	return &admin, token, nil
//}
//
//func (l *login) Logout(c *gin.Context, token string) errcode.Error {
//	if err := l.ssoHandler.SetOfflineStatus(token); err != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), err.Error())
//		return l.loginErr.SystemError()
//	}
//	if err := l.ssoHandler.ResignUserToken(token); err != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), err.Error())
//		return l.loginErr.SystemError()
//	}
//	return nil
//}
//
//
//func (l *login) LogoutForAdmin(c *gin.Context, token string) errcode.Error {
//	if err := l.ssoHandler.ResignAdminToken(token); err != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), err.Error())
//		return l.loginErr.SystemError()
//	}
//	return nil
//}
//
//func (l *login) login(c *gin.Context, identifier string, password string, role int) (*logindata.User, string, errcode.Error) {
//	l.logger.Set(c, handler.Warn, "LoginDebug", l.loginErr.SystemError().Code(), "Login Debug")
//	uid, err := l.userRepo.GetUserID(&repository.GetUserParam{
//		Identifier: &identifier,
//		Password:   &password,
//	})
//	if err != nil {
//		//查無此人
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, "", l.loginErr.LoginFailure()
//		}
//		//不明原因錯誤
//		l.logger.Set(c, handler.Error, "UserRepo", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	var user logindata.User
//	if err := l.userRepo.GetUserEntity(uid, &user); err != nil {
//		l.logger.Set(c, handler.Error, "UserRepo", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	if err := l.userRepo.GetUserInfoEntity(uid, &user.Info); err != nil {
//		l.logger.Set(c, handler.Error, "UserRepo", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	// 驗證帳號是否被停權
//	if user.Status == 2 {
//		return nil, "", l.loginErr.LoginStatusFailure()
//	}
//	// 驗證登入身份
//	if user.Role != role {
//		return nil, "", l.loginErr.LoginRoleFailure()
//	}
//	// 生成 Token
//	token, e := l.ssoHandler.GenerateToken(uid, user.Role)
//	if e != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), e.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	// 更新在線狀態
//	if err := l.ssoHandler.RenewOnlineStatus(token); err != nil {
//		l.logger.Set(c, handler.Error, "SSOHandler", l.loginErr.SystemError().Code(), err.Error())
//		return nil, "", l.loginErr.SystemError()
//	}
//	return &user, token, nil
//}