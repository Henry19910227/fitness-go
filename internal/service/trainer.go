package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
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
	albumRepo repository.TrainerAlbum
	cerRepo repository.Certificate
	uploader  handler.Uploader
	resHandler handler.Resource
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewTrainer(trainerRepo repository.Trainer, albumRepo repository.TrainerAlbum, cerRepo repository.Certificate, uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Trainer {
	return &trainer{trainerRepo: trainerRepo, albumRepo: albumRepo, cerRepo: cerRepo, uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}


func (t *trainer) CreateTrainer(c *gin.Context, uid int64, param *dto.CreateTrainerParam) (*dto.Trainer, errcode.Error) {
	//檢查教練身份是否存在
	isExists, e := t.trainerIsExists(c, uid)
	if e != nil {
		return nil, e
	}
	if isExists {
		return nil, t.errHandler.DataAlreadyExists()
	}
	//創建教練身份
	if err := t.trainerRepo.CreateTrainer(uid); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	var trainer dto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	return &trainer, nil
}

func (t *trainer) GetTrainerInfo(c *gin.Context, uid int64) (*dto.Trainer, errcode.Error) {
	//獲取trainer資訊
	var result dto.Trainer
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

func (t *trainer) GetTrainerInfoByToken(c *gin.Context, token string) (*dto.Trainer, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.GetTrainerInfo(c, uid)
}

func (t *trainer) UpdateTrainer(c *gin.Context, uid int64, param *dto.UpdateTrainerParam) (*dto.Trainer, errcode.Error) {
	//檢查教練身份是否存在
	isExists, e := t.trainerIsExists(c, uid)
	if e != nil {
		return nil, e
	}
	if !isExists {
		if err := t.trainerRepo.CreateTrainer(uid); err != nil {
			return nil, t.errHandler.Set(c, "trainer repo", err)
		}
	}
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		Name: param.Name,
		Nickname: param.Nickname,
		Email: param.Email,
		Phone: param.Phone,
		Address: param.Address,
		Intro: param.Intro,
		Experience: param.Experience,
		Motto: param.Motto,
		FacebookURL: param.FacebookURL,
		InstagramURL: param.InstagramURL,
		YoutubeURL: param.YoutubeURL,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	var trainer dto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	return &trainer, nil
}

func (t *trainer) UploadTrainerAvatarByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerAvatar, errcode.Error) {
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
	return &dto.TrainerAvatar{Avatar: newImageNamed}, nil
}

func (t *trainer) UploadTrainerAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.TrainerAvatar, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.UploadTrainerAvatarByUID(c, uid, imageNamed, imageFile)
}

func (t *trainer) UploadCardFrontImageByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerCardFront, errcode.Error) {
	//上傳照片
	newImageNamed, err := t.uploader.UploadCardFrontImage(imageFile, imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//查詢教練資訊
	var trainer struct{ CardFrontImage string `gorm:"column:card_front_image"`}
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//修改教練資訊
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		CardFrontImage: &newImageNamed,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//刪除舊照片
	if len(trainer.CardFrontImage) > 0 {
		if err := t.resHandler.DeleteCardFrontImage(trainer.CardFrontImage); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	return &dto.TrainerCardFront{Image: newImageNamed}, nil
}

func (t *trainer) UploadCardBackImageByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerCardBack, errcode.Error) {
	//上傳照片
	newImageNamed, err := t.uploader.UploadCardBackImage(imageFile, imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//查詢教練資訊
	var trainer struct{ CardBackImage string `gorm:"column:card_back_image"`}
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//修改教練資訊
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		CardBackImage: &newImageNamed,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//刪除舊照片
	if len(trainer.CardBackImage) > 0 {
		if err := t.resHandler.DeleteCardBackImage(trainer.CardBackImage); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	return &dto.TrainerCardBack{Image: newImageNamed}, nil
}

func (t *trainer) UploadAlbumPhoto(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerAlbumPhotoResult, errcode.Error) {
	//上傳照片
	newImageNamed, err := t.uploader.UploadTrainerAlbumPhoto(imageFile, imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//修改教練資訊
	if err := t.albumRepo.CreateAlbumPhoto(uid, newImageNamed); err != nil {
		return nil, t.errHandler.Set(c, "album repo", err)
	}
	return &dto.TrainerAlbumPhotoResult{Photo: newImageNamed}, nil
}

func (t *trainer) DeleteAlbumPhoto(c *gin.Context, photoID int64) errcode.Error {
	var albumPhoto dto.TrainerAlbumPhoto
	err := t.albumRepo.FindAlbumPhotoByID(photoID, &albumPhoto)
	if err != nil {
		return t.errHandler.Set(c, "album repo", err)
	}
	if err := t.albumRepo.DeleteAlbumPhotoByID(albumPhoto.ID); err != nil {
		return t.errHandler.Set(c, "album repo", err)
	}
	if err := t.resHandler.DeleteTrainerAlbumPhoto(albumPhoto.Photo); err != nil {
		t.errHandler.Set(c, "resource handler", err)
	}
	return nil
}

func (t *trainer) CreateCertificate(c *gin.Context, uid int64, name, imageNamed string, imageFile multipart.File) (*dto.Certificate, errcode.Error) {
	newImageNamed, err := t.uploader.UploadCertificateImage(imageFile, imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	cerID, err := t.cerRepo.CreateCertificate(uid, name, newImageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	var certificate dto.Certificate
	if err := t.cerRepo.FindCertificate(cerID, &certificate); err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	return &certificate, nil
}

func (t *trainer) DeleteCertificate(c *gin.Context, cerID int64) errcode.Error {
	var certificate dto.Certificate
	if err := t.cerRepo.FindCertificate(cerID, &certificate); err != nil {
		return t.errHandler.Set(c, "cer repo", err)
	}
	if err := t.cerRepo.DeleteCertificateByID(cerID); err != nil {
		return t.errHandler.Set(c, "cer repo", err)
	}
	if err := t.resHandler.DeleteCertificateImage(certificate.Image); err != nil {
		t.errHandler.Set(c, "resource handler", err)
	}
	return nil
}

func (t *trainer) trainerIsExists(c *gin.Context, uid int64) (bool, errcode.Error) {
	var trainer struct{
		UserID int64 `gorm:"column:user_id"`
		TrainerStatus int `gorm:"column:trainer_status"`
	}
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil{
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return false, t.errHandler.SystemError()
	}
	if trainer.UserID != 0 {
		return true, nil
	}
	return false, nil
}