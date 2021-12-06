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
	"strconv"
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
		Skill:              transformSkills(param.Skill),
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

func (t *trainer) UpdateTrainer(c *gin.Context, uid int64, param *dto.UpdateTrainerParam) (*dto.Trainer, errcode.Error) {
	//生成Avatar名稱
	var avatar *string
	if param.Avatar != nil {
		avatarImageNamed, err := t.uploader.GenerateNewImageName(param.Avatar.FileNamed)
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		avatar = &avatarImageNamed
	}
	//生成教練相簿照片名稱
	var createAlbumPhotoNames []string
	for _, file := range param.CreateAlbumPhotos {
		albumPhotoName, err := t.uploader.GenerateNewImageName(file.FileNamed);
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = albumPhotoName
		createAlbumPhotoNames = append(createAlbumPhotoNames, albumPhotoName)
	}
	//生成待更新證照照片名稱
	var updateCerImageNames []string
	for _, file := range param.UpdateCerImages {
		cerImageName, err := t.uploader.GenerateNewImageName(file.FileNamed);
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = cerImageName
		updateCerImageNames = append(updateCerImageNames, cerImageName)
	}
	//生成待新增證照照片名稱
	var createCerImageNames []string
	for _, file := range param.CreateCerImages {
		cerImageName, err := t.uploader.GenerateNewImageName(file.FileNamed);
		if err != nil {
			return nil, t.errHandler.Set(c, "uploader", err)
		}
		file.FileNamed = cerImageName
		createCerImageNames = append(createCerImageNames, cerImageName)
	}
	//轉換專長字串
	var skill = transformSkills(param.Skill)

	//查詢更新前的教練大頭照
	var oldTrainer struct{ Avatar string `gorm:"column:avatar"`}
	if err := t.trainerRepo.FindTrainerByUID(uid, &oldTrainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//查詢待刪除的教練相簿照片資料
	deleteAlbumItems := make([]*dto.TrainerAlbumPhoto, 0)
	if len(param.DeleteAlbumPhotosIDs) > 0 {
		if err := t.albumRepo.FindAlbumPhotosByIDs(param.DeleteAlbumPhotosIDs, &deleteAlbumItems); err != nil {
			return nil, t.errHandler.Set(c, "trainer album repo", err)
		}
		if len(deleteAlbumItems) != len(param.DeleteAlbumPhotosIDs) {
			return nil, t.errHandler.Set(c, "trainer album repo", errors.New(strconv.Itoa(errcode.DataNotFound)))
		}
	}
	//查詢待刪除的證照照片資料
	var deleteCerItems []*dto.Certificate
	if len(param.DeleteCerIDs) > 0 {
		if err := t.cerRepo.FindCertificatesByIDs(param.DeleteCerIDs, &deleteCerItems); err != nil {
			return nil, t.errHandler.Set(c, "trainer album repo", err)
		}
		if len(deleteCerItems) != len(param.DeleteCerIDs) {
			return nil, t.errHandler.Set(c, "trainer album repo", errors.New(strconv.Itoa(errcode.DataNotFound)))
		}
	}
	//查詢待更新前的證照照片資料
	var updateCerItems []*dto.Certificate
	if len(param.UpdateCerIDs) > 0 {
		if err := t.cerRepo.FindCertificatesByIDs(param.UpdateCerIDs, &updateCerItems); err != nil {
			return nil, t.errHandler.Set(c, "cer repo", err)
		}
		if len(updateCerItems) != len(param.UpdateCerIDs) {
			return nil, t.errHandler.Set(c, "cer repo", errors.New(strconv.Itoa(errcode.DataNotFound)))
		}
	}
	//修改資料
	if err := t.trainerRepo.UpdateTrainerByUID(uid, &model.UpdateTrainerParam{
		Nickname:              param.Nickname,
		Skill:                 &skill,
		Intro:                 param.Intro,
		Experience:            param.Experience,
		Motto:                 param.Motto,
		FacebookURL:           param.FacebookURL,
		InstagramURL:          param.InstagramURL,
		YoutubeURL:            param.YoutubeURL,
		Avatar:                avatar,
		DeleteAlbumPhotosIDs:  param.DeleteAlbumPhotosIDs,
		CreateAlbumPhotos:     createAlbumPhotoNames,
		DeleteCerIDs:          param.DeleteCerIDs,
		UpdateCerIDs:          param.UpdateCerIDs,
		UpdateCerImages:       updateCerImageNames,
		UpdateCerNames:        param.UpdateCerNames,
		CreateCerImages:       createCerImageNames,
		CreateCerNames:        param.CreateCerNames,
	}); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	//教練大頭照更新(刪除舊的 + 上傳新的)
	if param.Avatar != nil && avatar != nil{
		if err := t.resHandler.DeleteTrainerAvatar(oldTrainer.Avatar); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
		if err := t.uploader.UploadTrainerAvatar(param.Avatar.Data, *avatar); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//刪除指定教練相簿照片
	for _, item := range deleteAlbumItems {
		if err := t.resHandler.DeleteTrainerAlbumPhoto(item.Photo); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	//儲存指定新增的教練相簿照片
	for _, file := range param.CreateAlbumPhotos {
		if err := t.uploader.UploadTrainerAlbumPhoto(file.Data, file.FileNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//刪除指定刪除的證照照片
	for _, item := range deleteCerItems {
		if err := t.resHandler.DeleteCertificateImage(item.Image); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	//刪除指定更新前的證照照片
	for _, item := range updateCerItems {
		if err := t.resHandler.DeleteCertificateImage(item.Image); err != nil {
			t.errHandler.Set(c, "res handler", err)
		}
	}
	//儲存指定修改的證照照片
	for _, file := range param.UpdateCerImages {
		if err := t.uploader.UploadCertificateImage(file.Data, file.FileNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//儲存指定新增的證照照片
	for _, file := range param.CreateCerImages {
		if err := t.uploader.UploadCertificateImage(file.Data, file.FileNamed); err != nil {
			t.errHandler.Set(c, "uploader", err)
		}
	}
	//查詢更新完成資訊
	var trainer dto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	if err := t.albumRepo.FindAlbumPhotosByUID(uid, &trainer.TrainerAlbumPhotos); err != nil {
		return nil, t.errHandler.Set(c, "trainer album repo", err)
	}
	if err := t.cerRepo.FindCertificatesByUID(uid, &trainer.Certificates); err != nil {
		return nil, t.errHandler.Set(c, "cer repo", err)
	}
	return &trainer, nil
}

func (t *trainer) GetTrainer(c *gin.Context, uid int64) (*dto.Trainer, errcode.Error) {
	var trainer dto.Trainer
	if err := t.trainerRepo.FindTrainerByUID(uid, &trainer); err != nil {
		return nil, t.errHandler.Set(c, "trainer repo", err)
	}
	if trainer.UserID == 0 {
		return nil, t.errHandler.Set(c, "trainer repo", errors.New(strconv.Itoa(errcode.DataNotFound)))
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

func (t *trainer) GetTrainerAlbumPhotoCount(c *gin.Context, uid int64) (int, errcode.Error) {
	photos := make([]*dto.TrainerAlbumPhoto, 0)
	if err := t.albumRepo.FindAlbumPhotosByUID(uid, &photos); err != nil {
		return 0, t.errHandler.Set(c, "album repo", err)
	}
	return len(photos), nil
}

func (t *trainer) GetCertificateCount(c *gin.Context, uid int64) (int, errcode.Error) {
	cers := make([]*dto.Certificate, 0)
	if err := t.cerRepo.FindCertificatesByUID(uid, &cers); err != nil {
		return 0, t.errHandler.Set(c, "cer repo", err)
	}
	return len(cers), nil
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

func transformSkills(skills []int) string {
	var value string
	for i, skill := range skills {
		value += strconv.Itoa(skill)
		if i != len(skills) - 1 {
			value += ","
		}
	}
	return value
}