package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/userdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type user struct {
	Base
	userRepo  repository.User
	trainerRepo repository.Trainer
	logger    handler.Logger
	sso       handler.SSO
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer,
	logger handler.Logger, sso handler.SSO,
	jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo,
		trainerRepo: trainerRepo,
		logger: logger,
		sso: sso,
		jwtTool: jwtTool,
		errHandler: errHandler}
}

func (u *user) UpdateUserByToken(c *gin.Context, token string, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.UpdateUserByUID(c, uid, param)
}

func (u *user) UpdateUserByUID(c *gin.Context, uid int64, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error) {
	//檢查暱稱是否重複
	if param.Nickname != nil {
		_, err := u.userRepo.FindUserIDByNickname(*param.Nickname)
		//該暱稱已存在
		if err == nil {
			return nil, u.errHandler.NicknameDuplicate()
		}
		//不明原因錯誤
		if !errors.Is(err, gorm.ErrRecordNotFound){
			u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
			return nil, u.errHandler.SystemError()
		}
	}
	//更新user
	if err := u.userRepo.UpdateUserByUID(uid, &model.UpdateUserParam{
		Nickname: param.Nickname,
		Sex: param.Sex,
		Birthday: param.Birthday,
		Height: param.Height,
		Weight: param.Weight,
		Experience: param.Experience,
		Target: param.Target,
	}); err != nil {
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//查找user
	user, err := u.GetUserByUID(c, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *user) GetUserByUID(c *gin.Context, uid int64) (*userdto.User, errcode.Error) {
	//查找user
	var user userdto.User
	if err := u.userRepo.FindUserByUID(uid, &user); err != nil {
		//查無此資料
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, u.errHandler.DataNotFound()
		}
		//不明原因錯誤
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//檢查是否創建過教練身份
	err := u.trainerRepo.FindTrainerByUID(user.ID, nil)
	if err != nil {
		//不明原因錯誤
		if !errors.Is(err, gorm.ErrRecordNotFound){
			u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
			return nil, u.errHandler.SystemError()
		}
	} else { //教練身份已存在
		user.IsTrainer = 1
	}
	return &user, nil
}

func (u *user) GetUserByToken(c *gin.Context, token string) (*userdto.User, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.GetUserByUID(c, uid)
}

func (u *user) CreateTrainer(c *gin.Context, uid int64, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
	//檢查教練身份是否存在
	isExists, e := u.trainerIsExists(c, uid)
	if e != nil {
		return nil, e
	}
	if isExists {
		return nil, u.errHandler.DataAlreadyExists()
	}
	//創建教練身份
	err := u.trainerRepo.CreateTrainer(uid, &model.CreateTrainerParam{
		Name: param.Name,
		Nickname: param.Nickname,
		Phone: param.Phone,
		Email: param.Email,
	})
	if err != nil {
		//資料已存在
		if u.MysqlDuplicateEntry(err) {
			if strings.Contains(err.Error(), "nickname") {
				return nil, u.errHandler.Custom(9004, errors.New("重複的暱稱"))
			}
			return nil, u.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
		u.logger.Set(c, handler.Error, "Trainer Repo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	return &userdto.CreateTrainerResult{UserID: uid}, nil
}

func (u *user) CreateTrainerByToken(c *gin.Context, token string, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.CreateTrainer(c, uid, param)
}

func (u *user) GetTrainerInfo(c *gin.Context, uid int64) (*userdto.TrainerResult, errcode.Error) {
	//獲取trainer資訊
	var result userdto.TrainerResult
	if err := u.trainerRepo.FindTrainerByUID(uid, &result.Trainer); err != nil {
		//查無此資料
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, u.errHandler.DataNotFound()
		}
		//不明原因錯誤
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//產生 token
	token, err := u.sso.GenerateTrainerToken(uid)
	if err != nil {
		u.logger.Set(c, handler.Error, "sso", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	result.Token = token
	return &result, nil
}

func (u *user) GetTrainerInfoByToken(c *gin.Context, token string) (*userdto.TrainerResult, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.GetTrainerInfo(c, uid)
}

func (u *user) trainerIsExists(c *gin.Context, uid int64) (bool, errcode.Error) {
	err := u.trainerRepo.FindTrainerByUID(uid, nil)
	//教練身份已存在
	if err == nil {
		return true, nil
	}
	//不明原因錯誤
	if !errors.Is(err, gorm.ErrRecordNotFound){
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return false, u.errHandler.SystemError()
	}
	return false, nil
}

