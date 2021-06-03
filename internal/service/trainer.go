package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/trainerdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
)

type trainer struct {
	Base
	trainerRepo repository.Trainer
	uploader  handler.Uploader
	resHandler handler.Resource
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewTrainer(trainerRepo repository.Trainer, uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Trainer {
	return &trainer{trainerRepo: trainerRepo, uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}


func (t *trainer) CreateTrainer(c *gin.Context, uid int64, param *trainerdto.CreateTrainerParam) (*trainerdto.CreateTrainerResult, errcode.Error) {
	//檢查教練身份是否存在
	isExists, e := t.trainerIsExists(c, uid)
	if e != nil {
		return nil, e
	}
	if isExists {
		return nil, t.errHandler.DataAlreadyExists()
	}
	//創建教練身份
	err := t.trainerRepo.CreateTrainer(uid, &model.CreateTrainerParam{
		Name: param.Name,
		Nickname: param.Nickname,
		Phone: param.Phone,
		Email: param.Email,
	})
	if err != nil {
		//資料已存在
		if t.MysqlDuplicateEntry(err) {
			if strings.Contains(err.Error(), "nickname") {
				return nil, t.errHandler.Custom(9004, errors.New("重複的暱稱"))
			}
			return nil, t.errHandler.DataAlreadyExists()
		}
		//不明原因錯誤
		t.logger.Set(c, handler.Error, "Trainer Repo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	return &trainerdto.CreateTrainerResult{UserID: uid}, nil
}

func (t *trainer) CreateTrainerByToken(c *gin.Context, token string, param *trainerdto.CreateTrainerParam) (*trainerdto.CreateTrainerResult, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.CreateTrainer(c, uid, param)
}

func (t *trainer) GetTrainerInfo(c *gin.Context, uid int64) (*trainerdto.Trainer, errcode.Error) {
	//獲取trainer資訊
	var result trainerdto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &result); err != nil {
		//查無此資料
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, t.errHandler.DataNotFound()
		}
		//不明原因錯誤
		t.logger.Set(c, handler.Error, "UserRepo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	return &result, nil
}

func (t *trainer) GetTrainerInfoByToken(c *gin.Context, token string) (*trainerdto.Trainer, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.GetTrainerInfo(c, uid)
}

func (t *trainer) UploadTrainerAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*trainerdto.Avatar, errcode.Error) {
	//上傳照片
	newImageNamed, err := t.uploader.UploadTrainerAvatar(imageFile, imageNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, t.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, t.errHandler.FileSizeError()
		}
		t.logger.Set(c, handler.Error, "Resource Handler", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	//查詢教練資訊
	var trainer struct{ Avatar string `gorm:"column:avatar"`}
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	//修改教練資訊
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		Avatar: &newImageNamed,
	}); err != nil {
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	//刪除舊照片
	if len(trainer.Avatar) > 0 {
		if err := t.resHandler.DeleteTrainerAvatar(trainer.Avatar); err != nil {
			t.logger.Set(c, handler.Error, "ResHandler", t.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &trainerdto.Avatar{Avatar: newImageNamed}, nil
}

func (t *trainer) UploadTrainerAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*trainerdto.Avatar, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.UploadTrainerAvatarByUID(c, uid, imageNamed, imageFile)
}

func (t *trainer) trainerIsExists(c *gin.Context, uid int64) (bool, errcode.Error) {
	err := t.trainerRepo.FindTrainerByUID(uid, nil)
	//教練身份已存在
	if err == nil {
		return true, nil
	}
	//不明原因錯誤
	if !errors.Is(err, gorm.ErrRecordNotFound){
		t.logger.Set(c, handler.Error, "UserRepo", t.errHandler.SystemError().Code(), err.Error())
		return false, t.errHandler.SystemError()
	}
	return false, nil
}