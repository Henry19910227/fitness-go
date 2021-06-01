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

type trainer struct {
	Base
	trainerRepo repository.Trainer
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewTrainer(trainerRepo repository.Trainer, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Trainer {
	return &trainer{trainerRepo: trainerRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}


func (t *trainer) CreateTrainer(c *gin.Context, uid int64, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
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
	return &userdto.CreateTrainerResult{UserID: uid}, nil
}

func (t *trainer) CreateTrainerByToken(c *gin.Context, token string, param *userdto.CreateTrainerParam) (*userdto.CreateTrainerResult, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.CreateTrainer(c, uid, param)
}

func (t *trainer) GetTrainerInfo(c *gin.Context, uid int64) (*userdto.Trainer, errcode.Error) {
	//獲取trainer資訊
	var result userdto.Trainer
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

func (t *trainer) GetTrainerInfoByToken(c *gin.Context, token string) (*userdto.Trainer, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.GetTrainerInfo(c, uid)
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