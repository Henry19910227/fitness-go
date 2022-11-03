package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v2/field/trainer/required"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	fileModel "github.com/Henry19910227/fitness-go/internal/v2/model/file"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/trainer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type controller struct {
	resolver trainer.Resolver
}

func New(resolver trainer.Resolver) Controller {
	return &controller{resolver: resolver}
}

// CreateTrainer 創建教練
// @Summary 創建教練
// @Description 查看教練大頭照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/avatar/{圖片名} | 查看身分證正面照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/card_front_image/{圖片名} | 查看身分證背面照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/card_back_image/{圖片名} | 查看教練相簿照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/album/{圖片名} |  查看證照照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/certificate/{圖片名} |  查看銀行帳戶照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/account_image/{圖片名}
// @Tags 教練個人_v2
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param name formData string true "教練本名"
// @Param nickname formData string true "教練暱稱"
// @Param skill formData string true "專長-需按照順序排列(範例:1,3,5)(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
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
// @Success 200 {object} trainer.APICreateTrainerOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer [POST]
func (c *controller) CreateTrainer(ctx *gin.Context) {
	var input model.APICreateTrainerInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	// 獲取身分證正面照
	file, fileHeader, err := ctx.Request.FormFile("card_front_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input.CartFontImage.Named = fileHeader.Filename
	input.CartFontImage.Data = file
	// 獲取身分證背面照
	file, fileHeader, err = ctx.Request.FormFile("card_back_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input.CartBackImage.Named = fileHeader.Filename
	input.CartBackImage.Data = file
	// 獲取教練形象照
	file, fileHeader, err = ctx.Request.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input.Avatar.Named = fileHeader.Filename
	input.Avatar.Data = file
	// 獲取教練相簿照片
	fileDatas := ctx.Request.MultipartForm.File["trainer_album_photos"]
	files := make([]*fileModel.Input, 0)
	for _, fileData := range fileDatas {
		data, _ := fileData.Open()
		f := fileModel.Input{}
		f.Named = fileData.Filename
		f.Data = data
		files = append(files, &f)
	}
	input.TrainerAlbumPhotos = files
	// 獲取教練證照照片
	fileDatas = ctx.Request.MultipartForm.File["certificate_images"]
	files = make([]*fileModel.Input, 0)
	for _, fileData := range fileDatas {
		data, _ := fileData.Open()
		f := fileModel.Input{}
		f.Named = fileData.Filename
		f.Data = data
		files = append(files, &f)
	}
	input.CertificateImages = files
	// 獲取銀行帳戶照片
	file, fileHeader, err = ctx.Request.FormFile("account_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input.AccountImage.Named = fileHeader.Filename
	input.AccountImage.Data = file

	output := c.resolver.APICreateTrainer(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateTrainer 編輯教練
// @Summary 編輯教練
// @Description 查看教練大頭照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/avatar/{圖片名} | 查看身分證正面照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/card_front_image/{圖片名} | 查看身分證背面照 : https://www.fitopia-hub.tk/api/v2/resource/trainer/card_back_image/{圖片名} | 查看教練相簿照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/album/{圖片名} |  查看證照照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/certificate/{圖片名} |  查看銀行帳戶照片 : https://www.fitopia-hub.tk/api/v2/resource/trainer/account_image/{圖片名}
// @Tags 教練個人_v2
// @Accept mpfd
// @Produce json
// @Security fitness_token
// @Param nickname formData string false "教練暱稱"
// @Param skill formData string false "專長-需按照順序排列(範例:1,3,5)(1:功能性訓練/2:減脂/3:增肌/4:健美規劃/5:運動項目訓練/6:TRX/7:重量訓練/8:筋膜放鬆/9:瑜珈/10:體態雕塑/11:減重/12:心肺訓練/13:肌力訓練/14:其他)"
// @Param intro formData string false "教練介紹 (1~400字元)"
// @Param experience formData string false "年資 (0~40年)"
// @Param motto formData string false "座右銘 (1~100字元)"
// @Param facebook_url formData string false "臉書連結"
// @Param instagram_url formData string false "instagram連結"
// @Param youtube_url formData string false "youtube連結"
// @Param avatar formData file false "教練形象照"
// @Param delete_trainer_album_photos_id formData []int64 false "刪除的相簿照片id"
// @Param create_trainer_album_photos formData file false "新增的教練相簿照片(可一次新增多張)"
// @Param delete_certificate_id formData []int64 false "刪除的證照照片id"
// @Param update_certificate_id formData []int64 false "更新的證照照片id(可一次更新多個id)"
// @Param update_certificate_images formData file false "更新的證照照片(需與待更新的證照照片id數量相同)"
// @Param update_certificate_names formData []string false "更新的證照名稱(需與待更新的證照照片id數量相同)"
// @Param create_certificate_images formData file false "新增的證照照片(可一次新增多張)"
// @Param create_certificate_names formData []string false "新增的證照名稱(需與待新增的證照照片數量相同)"
// @Success 200 {object} trainer.APIUpdateTrainerOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer [PATCH]
func (c *controller) UpdateTrainer(ctx *gin.Context) {
	var input model.APIUpdateTrainerInput
	input.UserID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBind(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	// 獲取教練形象照
	file, fileHeader, _ := ctx.Request.FormFile("avatar")
	if file != nil {
		avatarFile := fileModel.Input{}
		avatarFile.Data = file
		avatarFile.Named = fileHeader.Filename
		input.Avatar = &avatarFile
	}
	// 獲取新增的教練相簿照片
	createAlbumDatas := ctx.Request.MultipartForm.File["create_trainer_album_photos"]
	createAlbumPhotos := make([]*fileModel.Input, 0)
	for _, fileData := range createAlbumDatas {
		data, _ := fileData.Open()
		f := fileModel.Input{}
		f.Named = fileData.Filename
		f.Data = data
		createAlbumPhotos = append(createAlbumPhotos, &f)
	}
	input.CreateAlbumPhotos = createAlbumPhotos
	// 獲取新增的證照照片
	createCertificateDatas := ctx.Request.MultipartForm.File["create_certificate_images"]
	createCertificateImages := make([]*fileModel.Input, 0)
	for _, fileData := range createCertificateDatas {
		data, _ := fileData.Open()
		f := fileModel.Input{}
		f.Named = fileData.Filename
		f.Data = data
		createCertificateImages = append(createCertificateImages, &f)
	}
	input.CreateCertificateImages = createCertificateImages
	// 獲取更新的證照照片
	updateCertificateDatas := ctx.Request.MultipartForm.File["update_certificate_images"]
	updateCertificateImages := make([]*fileModel.Input, 0)
	for _, fileData := range updateCertificateDatas {
		data, _ := fileData.Open()
		f := fileModel.Input{}
		f.Named = fileData.Filename
		f.Data = data
		updateCertificateImages = append(updateCertificateImages, &f)
	}
	input.UpdateCertificateImages = updateCertificateImages

	output := c.resolver.APIUpdateTrainer(ctx.MustGet("tx").(*gorm.DB), &input)
	ctx.JSON(http.StatusOK, output)
}

// GetTrainerProfile 獲取教練個人資訊
// @Summary 獲取教練個人資訊
// @Description 獲取教練個人資訊
// @Tags 教練個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} trainer.APIGetTrainerProfileOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/trainer/profile [GET]
func (c *controller) GetTrainerProfile(ctx *gin.Context) {
	var input model.APIGetTrainerProfileInput
	input.UserID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetTrainerProfile(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreTrainer 獲取商店教練詳細資訊
// @Summary 獲取商店教練詳細資訊
// @Description 獲取商店教練詳細資訊
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param user_id path int64 true "教練ID"
// @Success 200 {object} trainer.APIGetStoreTrainerOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/trainer/{user_id} [GET]
func (c *controller) GetStoreTrainer(ctx *gin.Context) {
	var input model.APIGetStoreTrainerInput
	if err := ctx.ShouldBindUri(&input.Uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreTrainer(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetStoreTrainers 獲取商店教練列表
// @Summary 獲取商店教練列表
// @Description 獲取商店教練列表
// @Tags 商店_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param order_field query string false "排序類型(latest:最新/popular:熱門)-單選"
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} trainer.APIGetStoreTrainersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/store/trainers [GET]
func (c *controller) GetStoreTrainers(ctx *gin.Context) {
	var input model.APIGetStoreTrainersInput
	if err := ctx.ShouldBindQuery(&input.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetStoreTrainers(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetFavoriteTrainers 獲取教練收藏列表
// @Summary 獲取教練收藏列表
// @Description 獲取教練收藏列表
// @Tags 收藏_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param page query int true "頁數(從第一頁開始)"
// @Param size query int true "筆數"
// @Success 200 {object} trainer.APIGetFavoriteTrainersOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/favorite/trainers [GET]
func (c *controller) GetFavoriteTrainers(ctx *gin.Context) {
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.JSON(http.StatusBadRequest, baseModel.InvalidToken())
		return
	}
	input := model.APIGetFavoriteTrainersInput{}
	input.UserID = uid.(int64)
	if err := ctx.ShouldBindQuery(&input.Form); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIGetFavoriteTrainers(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateCMSTrainerAvatar 更新教練大頭照
// @Summary 更新教練大頭照
// @Description 查看教練大頭照 : {Base URL}/v2/resource/trainer/avatar/{Filename}
// @Tags CMS會員管理_v2
// @Security fitness_token
// @Accept mpfd
// @Param user_id path int64 true "教練id"
// @Param avatar formData file true "教練大頭照"
// @Produce json
// @Success 200 {object} trainer.APIUpdateCMSTrainerAvatarOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/cms/trainer/{user_id}/avatar [PATCH]
func (c *controller) UpdateCMSTrainerAvatar(ctx *gin.Context) {
	var uri struct {
		required.UserIDField
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	file, fileHeader, err := ctx.Request.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIUpdateCMSTrainerAvatarInput{}
	input.UserID = uri.UserID
	input.CoverNamed = fileHeader.Filename
	input.File = file
	output := c.resolver.APIUpdateCMSTrainerAvatar(&input)
	ctx.JSON(http.StatusOK, output)
}
