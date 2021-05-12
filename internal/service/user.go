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
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
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
