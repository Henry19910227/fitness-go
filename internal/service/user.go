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
	"mime/multipart"
	"strings"
)

type user struct {
	Base
	userRepo  repository.User
	trainerRepo repository.Trainer
	uploader  handler.Uploader
	resHandler handler.Resource
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewUser(userRepo repository.User, trainerRepo repository.Trainer,
	uploader  handler.Uploader, resHandler handler.Resource, logger handler.Logger,
	jwtTool tool.JWT, errHandler errcode.Handler) User {
	return &user{userRepo: userRepo,
		trainerRepo: trainerRepo,
		uploader: uploader,
		resHandler: resHandler,
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

func (u *user) UploadUserAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*userdto.Avatar, errcode.Error) {
	//上傳照片
	newImageNamed, err := u.uploader.UploadUserAvatar(imageFile, imageNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, u.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, u.errHandler.FileSizeError()
		}
		u.logger.Set(c, handler.Error, "Resource Handler", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//查詢用戶資訊
	var user struct{ Avatar string `gorm:"column:avatar"`}
	if err := u.userRepo.FindUserByUID(uid, &user); err != nil {
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//修改教練資訊
	if err := u.userRepo.UpdateUserByUID(uid, &model.UpdateUserParam{
		Avatar: &newImageNamed,
	}); err != nil {
		u.logger.Set(c, handler.Error, "UserRepo", u.errHandler.SystemError().Code(), err.Error())
		return nil, u.errHandler.SystemError()
	}
	//刪除舊照片
	if len(user.Avatar) > 0 {
		if err := u.resHandler.DeleteUserAvatar(user.Avatar); err != nil {
			u.logger.Set(c, handler.Error, "ResHandler", u.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &userdto.Avatar{Avatar: newImageNamed}, nil
}

func (u *user) UploadUserAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*userdto.Avatar, errcode.Error) {
	uid, err := u.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, u.errHandler.InvalidToken()
	}
	return u.UploadUserAvatarByUID(c, uid, imageNamed, imageFile)
}


