package controller

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Trainer struct {
	Base
	trainerService service.Trainer
	courseService service.Course
}

func NewTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, courseService service.Course, userMiddleware gin.HandlerFunc, userMidd midd.User)  {
	baseGroup.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	baseGroup.StaticFS("/resource/trainer/card_front_image", http.Dir("./volumes/storage/trainer/card_front_image"))
	baseGroup.StaticFS("/resource/trainer/card_back_image", http.Dir("./volumes/storage/trainer/card_back_image"))
	baseGroup.StaticFS("/resource/trainer/album", http.Dir("./volumes/storage/trainer/album"))
	baseGroup.StaticFS("/resource/trainer/certificate", http.Dir("./volumes/storage/trainer/certificate"))
	baseGroup.StaticFS("/resource/trainer/account_image", http.Dir("./volumes/storage/trainer/account_image"))
	trainer := &Trainer{trainerService: trainerService, courseService: courseService}
	trainerGroup := baseGroup.Group("/trainer")
	trainerGroup.Use(userMiddleware)
	trainerGroup.GET("/info", trainer.GetTrainerInfo)

	baseGroup.POST("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		userMidd.TrainerAlbumPhotoLimit(nil, trainer.GetTrainerAlbumPhotoCount, nil, 5),
		userMidd.CertificateLimit(nil, trainer.GetCertificateCount, nil, 5),
		trainer.CreateTrainer)

	baseGroup.GET("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing, global.TrainerRevoke}),
		trainer.GetTrainer)

	baseGroup.GET("/trainers",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.GetTrainers)

	baseGroup.GET("/trainer/:user_id",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.GetTrainerByUID)

	baseGroup.GET("/trainer/:user_id/course_products",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		trainer.GetTrainerCourseProducts)

	baseGroup.PATCH("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerReviewing, global.TrainerRevoke}),
		userMidd.TrainerAlbumPhotoLimit(trainer.trainerService.GetTrainerAlbumPhotoCount, trainer.GetCreateAlbumPhotoCount, trainer.GetDeleteAlbumPhotoCount, 5),
		userMidd.CertificateLimit(trainer.trainerService.GetCertificateCount, trainer.GetCreateCertificateCount, trainer.GetDeleteCertificateCount, 20),
		trainer.UpdateTrainer)
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
// @Param intro formData string true "教練介紹 (1~400字元)"
// @Param experience formData string true "年資 (0~40年)"
// @Param motto formData string false "座右銘 (1~100字元)"
// @Param facebook_url formData string false "臉書連結"
// @Param instagram_url formData string false "instagram連結"
// @Param youtube_url formData string false "youtube連結"
// @Param avatar formData file true "教練形象照"
// @Param card_front_image formData file true "身分證正面照片"
// @Param card_back_image formData file true "身分證背面照片"
// @Param trainer_album_photos formData file false "教練相簿照片(可一次傳多張)"
// @Param certificate_images formData file false "證照照片(可一次傳多張)"
// @Param certificate_names formData []string false "證照名稱(需與證照照片數量相同)"
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
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		trainerAlbumPhotos = append(trainerAlbumPhotos, file)
	}
	//獲取教練證照照片
	files = c.Request.MultipartForm.File["certificate_images"]
	var certificateImages []*dto.File
	for _, f := range files {
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		certificateImages = append(certificateImages, file)
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

// UpdateTrainer 編輯教練
// @Summary 編輯教練
// @Description 查看教練大頭照 : https://www.fitness-app.tk/api/v1/resource/trainer/avatar/{圖片名} | 查看身分證正面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_front_image/{圖片名} | 查看身分證背面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_back_image/{圖片名} | 查看教練相簿照片 : https://www.fitness-app.tk/api/v1/resource/trainer/album/{圖片名} |  查看證照照片 : https://www.fitness-app.tk/api/v1/resource/trainer/certificate/{圖片名} |  查看銀行帳戶照片 : https://www.fitness-app.tk/api/v1/resource/trainer/account_image/{圖片名}
// @Tags Trainer
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param nickname formData string false "教練暱稱"
// @Param skill formData []int false "專長-需按照順序排列(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param intro formData string false "教練介紹 (1~400字元)"
// @Param experience formData string false "年資 (0~40年)"
// @Param motto formData string false "座右銘 (1~100字元)"
// @Param facebook_url formData string false "臉書連結"
// @Param instagram_url formData string false "instagram連結"
// @Param youtube_url formData string false "youtube連結"
// @Param avatar formData file false "教練形象照"
// @Param delete_trainer_album_photos_id formData []int64 false "待刪除的相簿照片id"
// @Param create_trainer_album_photos formData file false "待新增的教練相簿照片(可一次新增多張)"
// @Param delete_certificate_id formData []int64 false "待刪除的證照照片id"
// @Param update_certificate_id formData []int64 false "待更新的證照照片id(可一次更新多個id)"
// @Param update_certificate_images formData file false "待更新的證照照片(需與待更新的證照照片id數量相同)"
// @Param update_certificate_names formData []string false "待更新的證照名稱(需與待更新的證照照片id數量相同)"
// @Param create_certificate_images formData file false "待新增的證照照片(可一次新增多張)"
// @Param create_certificate_names formData []string false "待新增的證照名稱(需與待新增的證照照片數量相同)"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [PATCH]
func (t *Trainer) UpdateTrainer(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var form validator.UpdateTrainerForm
	if err := c.ShouldBind(&form); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//獲取形象照
	file, fileHeader, _ := c.Request.FormFile("avatar")
	var avatar *dto.File
	if file != nil {
		avatar = &dto.File{
			FileNamed: fileHeader.Filename,
			Data: file,
		}
	}
	//獲取教練相簿照片
	files := c.Request.MultipartForm.File["create_album_photos"]
	var createAlbumPhotos []*dto.File
	for _, f := range files {
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		createAlbumPhotos = append(createAlbumPhotos, file)
	}
	//獲取指定更新照片
	files = c.Request.MultipartForm.File["update_certificate_images"]
	var updateCerImages []*dto.File
	for _, f := range files {
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		updateCerImages = append(updateCerImages, file)
	}
	//獲取指定新增照片
	files = c.Request.MultipartForm.File["create_certificate_images"]
	var createCerImages []*dto.File
	for _, f := range files {
		data, _ := f.Open()
		file := &dto.File{
			FileNamed: f.Filename,
			Data: data,
		}
		createCerImages = append(createCerImages, file)
	}
	result, errs := t.trainerService.UpdateTrainer(c, uid, &dto.UpdateTrainerParam{
		Nickname: form.Nickname,
		Skill: form.Skill,
		Intro: form.Intro,
		Experience: form.Experience,
		Motto: form.Motto,
		FacebookURL: form.FacebookURL,
		InstagramURL: form.InstagramURL,
		YoutubeURL: form.YoutubeURL,
		Avatar: avatar,
		DeleteAlbumPhotosIDs: form.DeleteAlbumPhotosIDs,
		CreateAlbumPhotos: createAlbumPhotos,
		DeleteCerIDs: form.DeleteCerIDs,
		UpdateCerIDs: form.UpdateCerIDs,
		UpdateCerImages: updateCerImages,
		UpdateCerNames: form.UpdateCerNames,
		CreateCerNames: form.CreateCerNames,
		CreateCerImages: createCerImages,
	})
	if errs != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, result, "update success!")
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

// GetTrainers 取得教練列表
// @Summary 取得教練列表
// @Description 取得教練列表
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_type query string false "排序類型(latest:最新/popular:熱門)"
// @Param page query int true "頁數"
// @Param size query int true "每頁筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.TrainerSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainers [GET]
func (t *Trainer) GetTrainers(c *gin.Context) {
	var query validator.GetTrainerSummariesQuery
	if err := c.ShouldBind(&query); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	trainers, page, err := t.trainerService.GetTrainerSummaries(c, dto.GetTrainerSummariesParam{
		OrderType: query.OrderType,
	}, *query.Page, *query.Size)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessPagingResponse(c, trainers, page, "success!")
}

// GetTrainerCourseProducts 取得教練的課表產品清單
// @Summary 取得教練的課表產品清單
// @Description 取得教練的課表產品清單
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練ID"
// @Param sale_type query int false "銷售類型(1:免費課表/2:訂閱課表/3:付費課表)-單選"
// @Param page query int true "頁數"
// @Param size query int true "每頁筆數"
// @Success 200 {object} model.SuccessPagingResult{data=[]dto.CourseProductSummary} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer/{user_id}/course_products [GET]
func (t *Trainer) GetTrainerCourseProducts(c *gin.Context) {
	var uri validator.TrainerIDUri
	var query validator.GetTrainerCourseProductsQuery
	var pagingQuery validator.PagingQuery
	if err := c.ShouldBindUri(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindQuery(&pagingQuery); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	saleTypes := make([]int, 0)
	if query.SaleType != nil {
		saleTypes = append(saleTypes, *query.SaleType)
	}
	courses, paging, err := t.courseService.GetCourseProductSummaries(c, &dto.GetCourseProductSummariesParam{
		UserID: &uri.TrainerID,
		SaleType: saleTypes,
	}, pagingQuery.Page, pagingQuery.Size)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessPagingResponse(c, courses, paging, "success!")
}

// GetTrainerByUID 取得指定教練資訊
// @Summary 取得指定教練資訊
// @Description 取得指定教練資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練id"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer/{user_id} [GET]
func (t *Trainer) GetTrainerByUID(c *gin.Context) {
	var uri validator.TrainerIDUri
	if err := c.ShouldBindUri(&uri); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	trainer, err := t.trainerService.GetTrainer(c, uri.TrainerID)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, trainer, "success!")
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

func (t *Trainer) GetTrainerAlbumPhotoCount(c *gin.Context) int {
	_ = c.Request.ParseMultipartForm(10000000)
	files := c.Request.MultipartForm.File["trainer_album_photos"]
	return len(files)
}

func (t *Trainer) GetCertificateCount(c *gin.Context) int {
	_ = c.Request.ParseMultipartForm(10000000)
	files := c.Request.MultipartForm.File["certificate_images"]
	return len(files)
}

func (t *Trainer) GetCreateAlbumPhotoCount(c *gin.Context) int {
	_ = c.Request.ParseMultipartForm(10000000)
	files := c.Request.MultipartForm.File["create_album_photos"]
	return len(files)
}

func (t *Trainer) GetDeleteAlbumPhotoCount(c *gin.Context) int {
	_ = c.Request.ParseForm()
	form := struct {
		DeleteAlbumPhotosIDs []int64 `form:"delete_trainer_album_photos_id"`
	}{}
	if err := c.ShouldBind(&form); err != nil {
		return 0
	}
	return len(form.DeleteAlbumPhotosIDs)
}

func (t *Trainer) GetCreateCertificateCount(c *gin.Context) int {
	_ = c.Request.ParseMultipartForm(10000000)
	files := c.Request.MultipartForm.File["create_certificate_images"]
	return len(files)
}

func (t *Trainer) GetDeleteCertificateCount(c *gin.Context) int {
	_ = c.Request.ParseForm()
	form := struct {
		DeleteCerIDs []int64 `form:"delete_certificate_id"`
	}{}
	if err := c.ShouldBind(&form); err != nil {
		return 0
	}
	return len(form.DeleteCerIDs)
}