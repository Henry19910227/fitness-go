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
)

type user struct {
	Base
	userRepo  repository.User
	trainerRepo repository.Trainer
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo, trainerRepo: trainerRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (u *user) UpdateUserByToken(c *gin.Context, token string, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.UpdateUserByUID(c, uid, param)
}

func (u *user) UpdateUserByUID(c *gin.Context, uid int64, param *userdto.UpdateUserParam) (*userdto.User, errcode.Error) {
	//更新user
	if err := u.userRepo.UpdateUserByUID(uid, &model.UpdateUserParam{
		//Email: param.Email,
		//Nickname: param.Nickname,
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
	return &user, nil
}

func (u *user) CreateTrainer(c *gin.Context, uid int64, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
	//查詢是否已存在資料
	_, err := u.trainerRepo.FindTrainerIDByUID(uid)
	//教練身份已存在
	if err == nil {
		return nil, u.errHandler.DataAlreadyExists()
	}
	//不明原因錯誤
	if !errors.Is(err, gorm.ErrRecordNotFound){
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//創建教練身份
	trainerID, err := u.trainerRepo.CreateTrainer(uid, &model.CreateTrainerParam{
		Name: param.Name,
		Nickname: param.Nickname,
		Phone: param.Phone,
		Email: param.Email,
	})
	if err != nil {
		u.logger.Set(c, handler.Error, "Trainer Repo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	return &userdto.CreateTrainerResult{TrainerID: trainerID}, nil
}

func (u *user) CreateTrainerByToken(c *gin.Context, token string, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.CreateTrainer(c, uid, param)
}
