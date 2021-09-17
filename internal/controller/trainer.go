package controller

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Trainer struct {
	Base
	trainerService service.Trainer
}

func NewTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, userMiddleware gin.HandlerFunc, userMidd midd.User)  {
	baseGroup.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	baseGroup.StaticFS("/resource/trainer/card_front_image", http.Dir("./volumes/storage/trainer/card_front_image"))
	baseGroup.StaticFS("/resource/trainer/card_back_image", http.Dir("./volumes/storage/trainer/card_back_image"))
	baseGroup.StaticFS("/resource/trainer/album", http.Dir("./volumes/storage/trainer/album"))
	baseGroup.StaticFS("/resource/trainer/certificate", http.Dir("./volumes/storage/trainer/certificate"))
	baseGroup.StaticFS("/resource/trainer/account_image", http.Dir("./volumes/storage/trainer/account_image"))
	trainer := &Trainer{trainerService: trainerService}
	trainerGroup := baseGroup.Group("/trainer")
	trainerGroup.Use(userMiddleware)
	trainerGroup.GET("/info", trainer.GetTrainerInfo)
	trainerGroup.POST("/avatar", trainer.UploadMyTrainerAvatar)

	baseGroup.POST("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		trainer.CreateTrainer)

	baseGroup.GET("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing, global.TrainerRevoke}),
		trainer.GetTrainer)

	baseGroup.PATCH("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.UpdateTrainer)

	baseGroup.POST("/card_front_image",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.UploadCardFrontImage)

	baseGroup.POST("/card_back_image",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.UploadCardBackImage)

	baseGroup.POST("/trainer_album_photo",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerAlbumPhotoLimit(5),
		trainer.UploadTrainerAlbumPhoto)

	baseGroup.DELETE("/trainer_album_photo/:photo_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerAlbumPhotoCreatorVerify(),
		trainer.DeleteTrainerAlbumPhoto)

	baseGroup.POST("/certificate",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.CreateCertificate)

	baseGroup.DELETE("/certificate/:certificate_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.CertificateCreatorVerify(),
		trainer.DeleteCertificate)
}

// CreateTrainer 創建教練
// @Summary 創建教練
// @Description 查看教練大頭照 : https://www.fitness-app.tk/api/v1/resource/trainer/avatar/{圖片名} | 查看身分證正面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_front_image/{圖片名} | 查看身分證背面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_back_image/{圖片名} | 查看教練相簿照片 : https://www.fitness-app.tk/api/v1/resource/trainer/album/{圖片名} |  查看證照照片 : https://www.fitness-app.tk/api/v1/resource/trainer/certificate/{圖片名} |  查看銀行帳戶照片 : https://www.fitness-app.tk/api/v1/resource/trainer/account_image/{圖片名}
// @Tags Trainer
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param name formData string true "教練本名"
// @Param nickname formData string true "教練暱稱"
// @Param skill formData []int true "專長(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param email formData string true "信箱"
// @Param phone formData string true "手機"
// @Param address formData string true "地址 (最大100字元)"
// @Param Intro formData string true "教練介紹 (1~400字元)"
// @Param experience formData string true "年資 (0~40年)"
// @Param motto formData string false "座右銘 (1~100字元)"
// @Param facebook_url formData string false "臉書連結"
// @Param instagram_url formData string false "instagram連結"
// @Param youtube_url formData string false "youtube連結"
// @Param avatar formData file true "教練形象照"
// @Param card_front_image formData file true "身分證正面照片"
// @Param card_back_image formData file true "身分證背面照片"
// @Param trainer_album_photos formData file false "教練相簿照片(可一次傳多張)"
// @Param certificate_images formData file true "證照照片(可一次傳多張)"
// @Param certificate_names formData []string true "證照名稱(需與證照照片數量相同)"
// @Param account_name formData string true "帳戶名稱"
// @Param account formData string true "帳戶"
// @Param account_image formData file true "帳戶照片"
// @Param branch formData string true "分行"
// @Param bank_code formData string true "銀行代碼"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [POST]
func (t *Trainer) CreateTrainer(c *gin.Context)  {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var form validator.CreateTrainerForm
	if err := c.ShouldBind(&form); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//獲取身分證正面照
	file, fileHeader, err := c.Request.FormFile("card_front_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, errors.New("需上傳card_front_image").Error())
		return
	}
	cardFrontImage := &dto.File{
		FileNamed: fileHeader.Filename,
		Data: file,
	}
	//獲取身分證背面照
	file, fileHeader, err = c.Request.FormFile("card_back_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, errors.New("需上傳card_back_image").Error())
		return
	}
	cardBackImage := &dto.File{
		FileNamed: fileHeader.Filename,
		Data: file,
	}
	//獲取形象照
	file, fileHeader, err = c.Request.FormFile("avatar")
	if err != nil {
		t.JSONValidatorErrorResponse(c, errors.New("需上傳avatar").Error())
		return
	}
	avatar := &dto.File{
		FileNamed: fileHeader.Filename,
		Data: file,
	}
	//獲取教練相簿照片
	files := c.Request.MultipartForm.File["trainer_album_photos"]
	var trainerAlbumPhotos []*dto.File
	for _, f := range files {
		file := &dto.File{
			FileNamed: f.Filename,
			Data: file,
		}
		trainerAlbumPhotos = append(trainerAlbumPhotos, file)
	}
	if len(files) > 5 {
		t.JSONValidatorErrorResponse(c, errors.New(strconv.Itoa(errcode.FileCountError)).Error())
		return
	}
	//獲取教練證照照片
	files = c.Request.MultipartForm.File["certificate_images"]
	var certificateImages []*dto.File
	for _, f := range files {
		file := &dto.File{
			FileNamed: f.Filename,
			Data: file,
		}
		certificateImages = append(certificateImages, file)
	}
	if len(files) == 0 {
		t.JSONValidatorErrorResponse(c, errors.New("至少上傳一張certificate_images").Error())
		return
	}
	if len(files) > 20 {
		t.JSONValidatorErrorResponse(c, errors.New(strconv.Itoa(errcode.FileCountError)).Error())
		return
	}
	if len(certificateImages) != len(form.CerNames) {
		t.JSONValidatorErrorResponse(c, errors.New("證照名稱與照片數量不一致").Error())
		return
	}
	//獲取銀行帳戶照片
	file, fileHeader, err = c.Request.FormFile("account_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, errors.New("需上傳account_image").Error())
		return
	}
	accountImage := &dto.File{
		FileNamed: fileHeader.Filename,
		Data: file,
	}
	//創建教練
	result, errs := t.trainerService.CreateTrainer(c, uid, &dto.CreateTrainerParam{
		Name:               form.Name,
		Nickname:           form.Nickname,
		Skill:              form.Skill,
		Email:              form.Email,
		Phone:              form.Phone,
		Address:            form.Address,
		Intro:              form.Intro,
		Experience:         form.Experience,
		Motto:              form.Motto,
		FacebookURL:        form.FacebookURL,
		InstagramURL:       form.InstagramURL,
		YoutubeURL:         form.YoutubeURL,
		Avatar:             avatar,
		CardFrontImage:     cardFrontImage,
		CardBackImage:      cardBackImage,
		TrainerAlbumPhotos: trainerAlbumPhotos,
		CertificateImages:  certificateImages,
		CertificateNames:   form.CerNames,
		AccountName:        form.AccountName,
		AccountImage:       accountImage,
		BankCode:           form.BankCode,
		Account:            form.Account,
		Branch:             form.Branch,
	})
	if errs != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, result, "create success!")
}

// GetTrainer 取得我的教練資訊
// @Summary 取得我的教練資訊
// @Description 取得我的教練資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [GET]
func (t *Trainer) GetTrainer(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	trainer, err := t.trainerService.GetTrainer(c, uid)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, trainer, "success!")
}

// UpdateTrainer 修改我的教練資訊
// @Summary 修改我的教練資訊
// @Description 修改我的教練資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.UpdateTrainerBody true "輸入欄位"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [PATCH]
func (t *Trainer) UpdateTrainer(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.UpdateTrainerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	trainer, err := t.trainerService.UpdateTrainer(c, uid, &dto.UpdateTrainerParam{
		Name: body.Name,
		Nickname: body.Nickname,
		Email: body.Email,
		Phone: body.Phone,
		Address: body.Address,
		Intro: body.Intro,
		Experience: body.Experience,
		Motto: body.Motto,
		FacebookURL: body.FacebookURL,
		InstagramURL: body.InstagramURL,
		YoutubeURL: body.YoutubeURL,
	})
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, trainer, "update success!")
}

func (t *Trainer) GetTrainerInfo(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := t.trainerService.GetTrainerInfoByToken(c, header.Token)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, result, "success!")
}

func (t *Trainer) UploadMyTrainerAvatar(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := t.trainerService.UploadTrainerAvatarByToken(c, header.Token, fileHeader.Filename, file)
	if e != nil {
		t.JSONErrorResponse(c, e)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}

func (t *Trainer) UploadCardFrontImage(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("card_front_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, er := t.trainerService.UploadCardFrontImageByUID(c, uid, fileHeader.Filename, file)
	if er != nil {
		t.JSONErrorResponse(c, er)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}

func (t *Trainer) UploadCardBackImage(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("card_back_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, errs := t.trainerService.UploadCardBackImageByUID(c, uid, fileHeader.Filename, file)
	if errs != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}

func (t *Trainer) UploadTrainerAlbumPhoto(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("trainer_album_photo")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, errs := t.trainerService.UploadAlbumPhoto(c, uid, fileHeader.Filename, file)
	if errs != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}

func (t *Trainer) DeleteTrainerAlbumPhoto(c *gin.Context) {
	var uri validator.TrainerAlbumPhotoIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := t.trainerService.DeleteAlbumPhoto(c, uri.PhotoID); err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, nil, "delete success!")
}

func (t *Trainer) CreateCertificate(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	var uri validator.CreateCertificateQuery
	if err := c.ShouldBind(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	certificate, errs := t.trainerService.CreateCertificate(c, uid, uri.Name, fileHeader.Filename, file)
	if err != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, certificate, "create success!")
}

func (t *Trainer) UpdateCertificate(c *gin.Context) {

}

func (t *Trainer) DeleteCertificate(c *gin.Context) {
	var uri validator.CertificateIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := t.trainerService.DeleteCertificate(c, uri.CerID); err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, nil, "delete success!")
}