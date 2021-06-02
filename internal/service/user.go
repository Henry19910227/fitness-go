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
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer,
	logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo,
		trainerRepo: trainerRepo,
		logger: logger,
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
		//資料已存在
		if u.MysqlDuplicateEntry(err) {
			if strings.Contains(err.Error(), "nickname") {
				return nil, u.errHandler.Custom(9004, errors.New("重複的暱稱"))
			}
			return nil, u.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
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


