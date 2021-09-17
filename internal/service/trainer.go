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
	//生成Avatar名稱
	avatarImageNamed, err := t.uploader.GenerateNewImageName(param.Avatar.FileNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	param.Avatar.FileNamed = avatarImageNamed

	//生成CardFrontImage名稱
	cardFrontImageNamed, err := t.uploader.GenerateNewImageName(param.CardFrontImage.FileNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	param.CardFrontImage.FileNamed = cardFrontImageNamed

	//生成CardBackImage名稱
	cardBackImageNamed, err := t.uploader.GenerateNewImageName(param.CardBackImage.FileNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	param.CardBackImage.FileNamed = cardBackImageNamed

	//生成AccountImage名稱
	accountImageNamed, err := t.uploader.GenerateNewImageName(param.AccountImage.FileNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	param.AccountImage.FileNamed = accountImageNamed

	//生成教練相簿照片名稱
	var albumPhotoNames []string
	for _, file := range param.TrainerAlbumPhotos {
		albumPhotoName, err := t.uploader.GenerateNewImageName(file.FileNamed);
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = albumPhotoName
		albumPhotoNames = append(albumPhotoNames, albumPhotoName)
	}
	//生成證照圖片名稱
	var cerImageNames []string
	for _, file := range param.CertificateImages {
		cerImageName, err := t.uploader.GenerateNewImageName(file.FileNamed);
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = cerImageName
		cerImageNames = append(cerImageNames, cerImageName)
	}
	//創建教練身份
	if err := t.trainerRepo.CreateTrainer(uid, &model.CreateTrainerParam{
		Name:               param.Name,
		Nickname:           param.Nickname,
		Email:              param.Email,
		Phone:              param.Phone,
		Address:            param.Address,
		Intro:              param.Intro,
		Experience:         param.Experience,
		Motto:              param.Motto,
		CardFrontImage:     param.CardFrontImage.FileNamed,
		CardBackImage:      param.CardBackImage.FileNamed,
		FacebookURL:        param.FacebookURL,
		InstagramURL:       param.InstagramURL,
		YoutubeURL:         param.YoutubeURL,
		TrainerAlbumPhotos: albumPhotoNames,
		Avatar:             param.Avatar.FileNamed,
		CertificateImages:  cerImageNames,
		CertificateNames:   param.CertificateNames,
		AccountName:        param.AccountName,
		Branch:             param.Branch,
		AccountImage:       param.AccountImage.FileNamed,
		BankCode:           param.BankCode,
		Account:            param.Account,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//儲存教練形象照
	err = t.uploader.UploadTrainerAvatar(param.Avatar.Data, param.Avatar.FileNamed)
	if err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	//儲存身分證正面
	if err := t.uploader.UploadCardFrontImage(param.CardFrontImage.Data, param.CardFrontImage.FileNamed); err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	//儲存身分證背面
	if err := t.uploader.UploadCardBackImage(param.CardBackImage.Data, param.CardBackImage.FileNamed); err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	//儲存銀行帳戶照片
	if err := t.uploader.UploadAccountImage(param.AccountImage.Data, param.AccountImage.FileNamed); err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	//儲存教練相簿照片
	for _, file := range param.TrainerAlbumPhotos {
		if err := t.uploader.UploadTrainerAlbumPhoto(file.Data, file.FileNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//儲存證照照片
	for _, file := range param.CertificateImages {
		if err := t.uploader.UploadCertificateImage(file.Data, file.FileNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//查詢並返回結果
	var trainer dto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	if err := t.albumRepo.FindAlbumPhotosByUID(uid, &trainer.TrainerAlbumPhotos); err != nil {
		return nil, t.errHandler.Set(c, "trainer album repo", err)
	}
	if err := t.cerRepo.FindCertificatesByUID(uid, &trainer.Certificates); err != nil {
		return nil, t.errHandler.Set(c, "trainer album repo", err)
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
		if err := t.trainerRepo.CreateTrainer(uid, nil); err != nil {
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
	//生成Avatar名稱
	avatarImageNamed, err := t.uploader.GenerateNewImageName(imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//查詢教練資訊
	var trainer struct{ Avatar string `gorm:"column:avatar"`}
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	//修改教練資訊
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		Avatar: &avatarImageNamed,
	}); err != nil {
		t.logger.Set(c, handler.Error, "TrainerRepo", t.errHandler.SystemError().Code(), err.Error())
		return nil, t.errHandler.SystemError()
	}
	//上傳教練形象照
	err = t.uploader.UploadTrainerAvatar(imageFile, avatarImageNamed)
	if err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	//刪除舊照片
	if len(trainer.Avatar) > 0 {
		if err := t.resHandler.DeleteTrainerAvatar(trainer.Avatar); err != nil {
			t.logger.Set(c, handler.Error, "ResHandler", t.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &dto.TrainerAvatar{Avatar: avatarImageNamed}, nil
}

func (t *trainer) UploadTrainerAvatarByToken(c *gin.Context, token string, imageNamed string, imageFile multipart.File) (*dto.TrainerAvatar, errcode.Error) {
	uid, err := t.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, t.errHandler.InvalidToken()
	}
	return t.UploadTrainerAvatarByUID(c, uid, imageNamed, imageFile)
}

func (t *trainer) UploadCardFrontImageByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerCardFront, errcode.Error) {
	//生成CardFrontImage名稱
	cardFrontImageNamed, err := t.uploader.GenerateNewImageName(imageNamed)
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
		CardFrontImage: &cardFrontImageNamed,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//刪除舊照片
	if len(trainer.CardFrontImage) > 0 {
		if err := t.resHandler.DeleteCardFrontImage(trainer.CardFrontImage); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	//上傳照片
	err = t.uploader.UploadCardFrontImage(imageFile, cardFrontImageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	return &dto.TrainerCardFront{Image: cardFrontImageNamed}, nil
}

func (t *trainer) UploadCardBackImageByUID(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerCardBack, errcode.Error) {
	//生成CardBackImage名稱
	cardBackImageNamed, err := t.uploader.GenerateNewImageName(imageNamed)
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
		CardBackImage: &cardBackImageNamed,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//上傳照片
	err = t.uploader.UploadCardBackImage(imageFile, cardBackImageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//刪除舊照片
	if len(trainer.CardBackImage) > 0 {
		if err := t.resHandler.DeleteCardBackImage(trainer.CardBackImage); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	return &dto.TrainerCardBack{Image: cardBackImageNamed}, nil
}

func (t *trainer) UploadAlbumPhoto(c *gin.Context, uid int64, imageNamed string, imageFile multipart.File) (*dto.TrainerAlbumPhotoResult, errcode.Error) {
	newImageNamed, err := t.uploader.GenerateNewImageName(imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	//修改教練資訊
	if err := t.albumRepo.CreateAlbumPhoto(uid, newImageNamed); err != nil {
		return nil, t.errHandler.Set(c, "album repo", err)
	}
	//上傳照片
	err = t.uploader.UploadTrainerAlbumPhoto(imageFile, imageNamed)
	if err != nil {
		t.errHandler.Set(c, "uploader", err)
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
	newImageNamed, err := t.uploader.GenerateNewImageName(imageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "uploader", err)
	}
	cerID, err := t.cerRepo.CreateCertificate(uid, name, newImageNamed)
	if err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	err = t.uploader.UploadCertificateImage(imageFile, imageNamed)
	if err != nil {
		t.errHandler.Set(c, "uploader", err)
	}
	var certificate dto.Certificate
	if err := t.cerRepo.FindCertificate(cerID, &certificate); err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	return &certificate, nil
}

func (t *trainer) UpdateCertificate(c *gin.Context, cerID int64, name *string, file *dto.File) (*dto.Certificate, errcode.Error) {
	var newImageNamed *string
	var oldImageNamed string
	if file != nil {
		var certificate dto.Certificate
		if err := t.cerRepo.FindCertificate(cerID, &certificate); err != nil {
			return nil, t.errHandler.Set(c, "cer repo", err)
		}
		imageNamed, err := t.uploader.GenerateNewImageName(file.FileNamed)
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		newImageNamed = &imageNamed
		oldImageNamed = certificate.Image
	}
	//更新證照
	if err := t.cerRepo.UpdateCertificate(cerID, name, newImageNamed); err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	//上傳圖片
	if file != nil && newImageNamed != nil {
		if err := t.uploader.UploadCertificateImage(file.Data, *newImageNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	if len(oldImageNamed) > 0 {
		if err := t.resHandler.DeleteCertificateImage(oldImageNamed); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
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