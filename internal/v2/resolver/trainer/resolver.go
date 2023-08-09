package trainer

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/fcm"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/logger"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/redis"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	accountModel "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	cardModel "github.com/Henry19910227/fitness-go/internal/v2/model/card"
	certModel "github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	courseModel "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	fcmModel "github.com/Henry19910227/fitness-go/internal/v2/model/fcm"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	preloadModel "github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer/api_get_cms_trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer/api_get_cms_trainers"
	"github.com/Henry19910227/fitness-go/internal/v2/model/trainer/api_update_cms_trainer"
	albumModel "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_album"
	trainerStatusLogModel "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_status_update_log"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/service/card"
	"github.com/Henry19910227/fitness-go/internal/v2/service/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_album"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_status_update_log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type resolver struct {
	trainerService          trainer.Service
	trainerAlbumService     trainer_album.Service
	trainerStatusLogService trainer_status_update_log.Service
	cardService             card.Service
	certService             certificate.Service
	bankAccountService      bank_account.Service
	courseService           course.Service
	avatarUploadTool        uploader.Tool
	albumUploadTool         uploader.Tool
	cardFrontUploadTool     uploader.Tool
	cardBackUploadTool      uploader.Tool
	certUploadTool          uploader.Tool
	accountUploadTool       uploader.Tool
	redisTool               redis.Tool
	fcmTool                 fcm.Tool
}

func New(trainerService trainer.Service, trainerAlbumService trainer_album.Service,
	trainerStatusLogService trainer_status_update_log.Service,
	cardService card.Service, certService certificate.Service,
	bankAccountService bank_account.Service, courseService course.Service,
	avatarUploadTool uploader.Tool, albumUploadTool uploader.Tool,
	cardFrontUploadTool uploader.Tool, cardBackUploadTool uploader.Tool,
	certUploadTool uploader.Tool, accountUploadTool uploader.Tool,
	redisTool redis.Tool, fcmTool fcm.Tool) Resolver {
	return &resolver{trainerService: trainerService, trainerAlbumService: trainerAlbumService,
		trainerStatusLogService: trainerStatusLogService,
		cardService:             cardService, certService: certService,
		bankAccountService: bankAccountService, courseService: courseService,
		avatarUploadTool: avatarUploadTool, albumUploadTool: albumUploadTool,
		cardFrontUploadTool: cardFrontUploadTool, cardBackUploadTool: cardBackUploadTool,
		certUploadTool: certUploadTool, accountUploadTool: accountUploadTool,
		redisTool: redisTool, fcmTool: fcmTool}
}

func (r *resolver) APICreateTrainer(tx *gorm.DB, input *model.APICreateTrainerInput) (output model.APICreateTrainerOutput) {
	// 輸入驗證
	if len(input.TrainerAlbumPhotos) > 5 {
		output.Set(code.BadRequest, "教練相簿照片上傳超過五張上限")
		return output
	}
	if len(input.CertificateImages) != len(input.Form.CertificateNames) {
		output.Set(code.BadRequest, "證照名稱與照片數量不一致")
		return output
	}
	// 權限驗證
	trainerListInput := model.ListInput{}
	trainerListInput.UserID = util.PointerInt64(input.UserID)
	trainerOutputs, _, err := r.trainerService.List(&trainerListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(trainerOutputs) > 0 {
		output.Set(code.BadRequest, "此帳號已經擁有教練身份")
		return output
	}
	// 儲存教練頭貼
	avatar, err := r.avatarUploadTool.Save(input.Avatar.Data, input.Avatar.Named)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存身分證正面照片
	cardFront, err := r.cardFrontUploadTool.Save(input.CartFontImage.Data, input.CartFontImage.Named)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存身分證背面照片
	cardBack, err := r.cardBackUploadTool.Save(input.CartBackImage.Data, input.CartBackImage.Named)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存銀行帳戶照片
	accountImage, err := r.accountUploadTool.Save(input.AccountImage.Data, input.AccountImage.Named)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 儲存教練相簿圖片
	albumTables := make([]*albumModel.Table, 0)
	for _, file := range input.TrainerAlbumPhotos {
		photo, _ := r.albumUploadTool.Save(file.Data, file.Named)
		table := albumModel.Table{}
		table.UserID = util.PointerInt64(input.UserID)
		table.Photo = util.PointerString(photo)
		albumTables = append(albumTables, &table)
	}
	// 儲存證照圖片
	certTables := make([]*certModel.Table, 0)
	for idx, file := range input.CertificateImages {
		image, _ := r.certUploadTool.Save(file.Data, file.Named)
		table := certModel.Table{}
		table.UserID = util.PointerInt64(input.UserID)
		table.Name = util.PointerString(input.Form.CertificateNames[idx])
		table.Image = util.PointerString(image)
		certTables = append(certTables, &table)
	}
	defer tx.Rollback()
	// 創建教練
	trainerTable := model.Table{}
	trainerTable.UserID = util.PointerInt64(input.UserID)
	trainerTable.Name = util.PointerString(input.Form.Name)
	trainerTable.Nickname = util.PointerString(input.Form.Nickname)
	trainerTable.Skill = util.PointerString(input.Form.Skill)
	trainerTable.Avatar = util.PointerString(avatar)
	trainerTable.TrainerStatus = util.PointerInt(model.Reviewing)
	trainerTable.TrainerLevel = util.PointerInt(1)
	trainerTable.Email = util.PointerString(input.Form.Email)
	trainerTable.Phone = util.PointerString(input.Form.Phone)
	trainerTable.Address = util.PointerString(input.Form.Address)
	trainerTable.Intro = util.PointerString(input.Form.Intro)
	trainerTable.Experience = util.PointerInt(input.Form.Experience)
	trainerTable.Motto = input.Form.Motto
	trainerTable.FacebookURL = input.Form.FacebookURL
	trainerTable.InstagramURL = input.Form.InstagramURL
	trainerTable.YoutubeURL = input.Form.YoutubeURL
	if err := r.trainerService.Tx(tx).Create(&trainerTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建教練身分證
	cardTable := cardModel.Table{}
	cardTable.UserID = util.PointerInt64(input.UserID)
	cardTable.CardID = util.PointerString("") //暫時不存身分證字號
	cardTable.FrontImage = util.PointerString(cardFront)
	cardTable.BackImage = util.PointerString(cardBack)
	if err := r.cardService.Tx(tx).Create(&cardTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建教練相簿圖片
	if err := r.trainerAlbumService.Tx(tx).Creates(albumTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建證照
	if err := r.certService.Tx(tx).Creates(certTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 創建銀行帳戶
	accountTable := accountModel.Table{}
	accountTable.UserID = util.PointerInt64(input.UserID)
	accountTable.Account = util.PointerString(input.Form.Account)
	accountTable.AccountName = util.PointerString(input.Form.AccountName)
	accountTable.AccountImage = util.PointerString(accountImage)
	accountTable.BankCode = util.PointerString(input.Form.BankCode)
	accountTable.Branch = util.PointerString(input.Form.Branch)
	if err := r.bankAccountService.Tx(tx).Create(&accountTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// Parser output
	findTrainerInput := model.FindInput{}
	findTrainerInput.UserID = util.PointerInt64(input.UserID)
	findTrainerInput.Preloads = []*preload.Preload{{Field: "TrainerStatistic"}, {Field: "Certificates"}, {Field: "TrainerAlbums"}}
	trainerOutput, err := r.trainerService.Find(&findTrainerInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APICreateTrainerData{}
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateTrainer(tx *gorm.DB, input *model.APIUpdateTrainerInput) (output model.APIUpdateTrainerOutput) {
	// 查詢教練資訊
	findInput := model.FindInput{}
	findInput.Preloads = []*preload.Preload{{Field: "Certificates"}, {Field: "TrainerAlbums"}}
	trainerOutput, err := r.trainerService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 驗證教練相簿刪除限制
	albumListInput := albumModel.ListInput{}
	albumListInput.Wheres = []*whereModel.Where{
		{Query: "trainer_albums.id IN (?)", Args: []interface{}{input.Form.DeleteAlbumPhotosIDs}},
	}
	deleteAlbumOutputs, _, err := r.trainerAlbumService.List(&albumListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, albumOutput := range deleteAlbumOutputs {
		if util.OnNilJustReturnInt64(albumOutput.UserID, 0) != input.UserID {
			output.Set(code.BadRequest, "非此教練相簿照片擁有者，無法刪除資源")
			return output
		}
	}
	if len(deleteAlbumOutputs) != len(input.Form.DeleteAlbumPhotosIDs) {
		output.Set(code.BadRequest, "查無某些教練相簿照片，無法刪除資源")
		return output
	}
	// 驗證教練相簿新增限制
	if len(trainerOutput.TrainerAlbums)-len(input.Form.DeleteAlbumPhotosIDs)+len(input.CreateAlbumPhotos) > 5 {
		output.Set(code.BadRequest, "教練相簿數量超過上限，無法新增資源")
		return output
	}
	// 驗證證照刪除限制
	certListInput := certModel.ListInput{}
	certListInput.Wheres = []*whereModel.Where{
		{Query: "certificates.id IN (?)", Args: []interface{}{input.Form.DeleteCertificateIDs}},
	}
	deleteCertOutputs, _, err := r.certService.List(&certListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, certOutput := range deleteCertOutputs {
		if util.OnNilJustReturnInt64(certOutput.UserID, 0) != input.UserID {
			output.Set(code.BadRequest, "非此證照照片擁有者，無法刪除資源")
			return output
		}
	}
	if len(deleteCertOutputs) != len(input.Form.DeleteCertificateIDs) {
		output.Set(code.BadRequest, "查無某些證照照片，無法刪除資源")
		return output
	}
	// 驗證證照更新限制
	certListInput = certModel.ListInput{}
	certListInput.Wheres = []*whereModel.Where{
		{Query: "certificates.id IN (?)", Args: []interface{}{input.Form.UpdateCertificateIDs}},
	}
	updateCertOutputs, _, err := r.certService.List(&certListInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	for _, certOutput := range updateCertOutputs {
		if util.OnNilJustReturnInt64(certOutput.UserID, 0) != input.UserID {
			output.Set(code.BadRequest, "非此證照照片擁有者，無法更新資源")
			return output
		}
	}
	if len(updateCertOutputs) != len(input.Form.UpdateCertificateIDs) {
		output.Set(code.BadRequest, "查無某些證照照片，無法更新資源")
		return output
	}
	if len(input.Form.UpdateCertificateIDs) != len(input.UpdateCertificateImages) || len(input.Form.UpdateCertificateIDs) != len(input.Form.UpdateCertificateNames) {
		output.Set(code.BadRequest, "更新證照的照片或名稱與證照ID數量不相等，無法更新資源")
		return output
	}
	// 驗證證照新增限制
	if len(input.Form.CreateCertificateNames) != len(input.CreateCertificateImages) {
		output.Set(code.BadRequest, "新增證照的照片與名稱數量不相等，無法更新資源")
		return output
	}
	defer tx.Rollback()
	// 刪除教練相簿照片
	albumDeleteInputs := make([]*albumModel.DeleteInput, 0)
	for _, id := range input.Form.DeleteAlbumPhotosIDs {
		deleteInput := albumModel.DeleteInput{}
		deleteInput.ID = util.PointerInt64(id)
		albumDeleteInputs = append(albumDeleteInputs, &deleteInput)
	}
	if err := r.trainerAlbumService.Tx(tx).Deletes(albumDeleteInputs); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 刪除證照照片
	certDeleteInputs := make([]*certModel.DeleteInput, 0)
	for _, id := range input.Form.DeleteCertificateIDs {
		deleteInput := certModel.DeleteInput{}
		deleteInput.ID = util.PointerInt64(id)
		certDeleteInputs = append(certDeleteInputs, &deleteInput)
	}
	if err := r.certService.Tx(tx).Deletes(certDeleteInputs); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新教練資訊
	trainerTable := model.Table{}
	trainerTable.UserID = util.PointerInt64(input.UserID)
	trainerTable.Nickname = input.Form.Nickname
	trainerTable.Skill = input.Form.Skill
	trainerTable.Intro = input.Form.Intro
	trainerTable.Experience = input.Form.Experience
	trainerTable.Motto = input.Form.Motto
	trainerTable.FacebookURL = input.Form.FacebookURL
	trainerTable.InstagramURL = input.Form.InstagramURL
	trainerTable.YoutubeURL = input.Form.YoutubeURL
	if input.Avatar != nil {
		avatar, _ := r.avatarUploadTool.Save(input.Avatar.Data, input.Avatar.Named)
		trainerTable.Avatar = util.PointerString(avatar)
	}
	if err := r.trainerService.Tx(tx).Update(&trainerTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 更新證照照片
	certTables := make([]*certModel.Table, 0)
	for idx, id := range input.Form.UpdateCertificateIDs {
		file := input.UpdateCertificateImages[idx]
		name := input.Form.UpdateCertificateNames[idx]
		imageNamed, _ := r.certUploadTool.Save(file.Data, file.Named)
		certTable := certModel.Table{}
		certTable.ID = util.PointerInt64(id)
		certTable.Name = util.PointerString(name)
		certTable.Image = util.PointerString(imageNamed)
		certTables = append(certTables, &certTable)
	}
	if err := r.certService.Tx(tx).Updates(certTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 新增證照照片
	certTables = make([]*certModel.Table, 0)
	for idx, name := range input.Form.CreateCertificateNames {
		file := input.CreateCertificateImages[idx]
		imageNamed, _ := r.certUploadTool.Save(file.Data, file.Named)
		certTable := certModel.Table{}
		certTable.UserID = util.PointerInt64(input.UserID)
		certTable.Name = util.PointerString(name)
		certTable.Image = util.PointerString(imageNamed)
		certTables = append(certTables, &certTable)
	}
	if err := r.certService.Tx(tx).Creates(certTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 新增教練相簿照片
	albumTables := make([]*albumModel.Table, 0)
	for _, file := range input.CreateAlbumPhotos {
		imageNamed, _ := r.albumUploadTool.Save(file.Data, file.Named)
		albumTable := albumModel.Table{}
		albumTable.UserID = util.PointerInt64(input.UserID)
		albumTable.Photo = util.PointerString(imageNamed)
		albumTables = append(albumTables, &albumTable)
	}
	if err := r.trainerAlbumService.Tx(tx).Creates(albumTables); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	tx.Commit()
	// 刪除已經刪除的教練相簿檔案
	for _, albumOutput := range deleteAlbumOutputs {
		_ = r.albumUploadTool.Delete(util.OnNilJustReturnString(albumOutput.Photo, ""))
	}
	// 刪除已經刪除的證照檔案
	for _, certOutput := range deleteCertOutputs {
		_ = r.certUploadTool.Delete(util.OnNilJustReturnString(certOutput.Image, ""))
	}
	// 刪除已經更新的證照檔案
	for _, certOutput := range updateCertOutputs {
		_ = r.certUploadTool.Delete(util.OnNilJustReturnString(certOutput.Image, ""))
	}
	// Parser output
	findTrainerInput := model.FindInput{}
	findTrainerInput.UserID = util.PointerInt64(input.UserID)
	findTrainerInput.Preloads = []*preload.Preload{{Field: "TrainerStatistic"}, {Field: "Certificates"}, {Field: "TrainerAlbums"}}
	trainerOutput, err = r.trainerService.Find(&findTrainerInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIUpdateTrainerData{}
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Preloads = []*preload.Preload{{Field: "TrainerStatistic"}, {Field: "Certificates"}, {Field: "TrainerAlbums"}}
	listInput.Size = util.PointerInt(1)
	listInput.Page = util.PointerInt(1)
	listInput.OrderField = "create_at"
	listInput.OrderType = order_by.DESC
	trainerOutputs, _, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	if len(trainerOutputs) == 0 {
		output.Set(code.DataNotFound, "查無資料")
		return output
	}
	data := model.APIGetTrainerProfileData{}
	if err := util.Parser(trainerOutputs[0], &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreTrainer(input *model.APIGetStoreTrainerInput) (output model.APIGetStoreTrainerOutput) {
	findInput := model.FindInput{}
	findInput.UserID = util.PointerInt64(input.Uri.UserID)
	findInput.Preloads = []*preload.Preload{
		{Field: "User"},
		{Field: "TrainerStatistic"},
		{Field: "Certificates"},
		{Field: "TrainerAlbums"},
	}
	trainerOutput, err := r.trainerService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetStoreTrainerData{}
	data.IsDeleted = trainerOutput.UserOnSafe().IsDeleted
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIGetStoreTrainers(input *model.APIGetStoreTrainersInput) (output model.APIGetStoreTrainersOutput) {
	joins := make([]*joinModel.Join, 0)
	wheres := make([]*whereModel.Where, 0)
	orders := make([]*orderByModel.Order, 0)

	joins = append(joins, &joinModel.Join{Query: "INNER JOIN users ON users.id = trainers.user_id"})
	wheres = append(wheres, &whereModel.Where{Query: "users.is_deleted = ?", Args: []interface{}{0}})
	wheres = append(wheres, &whereModel.Where{Query: "trainers.trainer_status = ?", Args: []interface{}{model.Activity}})
	if input.Query.OrderField != nil {
		if *input.Query.OrderField == "latest" {
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainers.%s %s", "create_at", order_by.DESC)})
		}
		if *input.Query.OrderField == "popular" {
			joins = append(joins, &joinModel.Join{Query: "LEFT JOIN trainer_statistics ON trainers.user_id = trainer_statistics.user_id"})
			orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainer_statistics.%s %s", "student_count", order_by.DESC)})
		}
	} else {
		orders = append(orders, &orderByModel.Order{Value: fmt.Sprintf("trainers.%s %s", "create_at", order_by.DESC)})
	}
	listInput := model.ListInput{}
	listInput.Wheres = wheres
	listInput.Joins = joins
	listInput.Orders = orders
	listInput.Size = input.Query.Size
	listInput.Page = input.Query.Page
	trainerOutputs, page, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	data := model.APIGetStoreTrainersData{}
	if err := util.Parser(trainerOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetFavoriteTrainers(input *model.APIGetFavoriteTrainersInput) (output model.APIGetFavoriteTrainersOutput) {
	// parser input
	listInput := model.ListInput{}
	listInput.Joins = []*joinModel.Join{
		{Query: "INNER JOIN favorite_trainers ON trainers.user_id = favorite_trainers.trainer_id"},
		{Query: "INNER JOIN users ON trainers.user_id = users.id"},
	}
	listInput.Wheres = []*whereModel.Where{
		{Query: "favorite_trainers.user_id = ?", Args: []interface{}{input.UserID}},
		{Query: "users.is_deleted = ?", Args: []interface{}{0}},
	}
	listInput.Orders = []*orderByModel.Order{
		{Value: fmt.Sprintf("favorite_trainers.%s %s", "create_at", order_by.DESC)},
	}
	if err := util.Parser(input.Form, &listInput); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 執行查詢
	results, page, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// parser output
	data := model.APIGetFavoriteTrainersData{}
	if err := util.Parser(results, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = data
	return output
}

func (r *resolver) APIGetCMSTrainers(input *api_get_cms_trainers.Input) (output api_get_cms_trainers.Output) {
	// 查詢列表教練資訊
	listInput := model.ListInput{}
	listInput.UserID = input.Query.UserID
	listInput.Nickname = input.Query.Nickname
	listInput.TrainerStatus = input.Query.TrainerStatus
	listInput.Page = input.Query.Page
	listInput.Size = input.Query.Size
	listInput.OrderType = orderByModel.DESC
	listInput.OrderField = "create_at"
	if input.Query.OrderType != nil {
		listInput.OrderType = *input.Query.OrderType
	}
	if input.Query.OrderField != nil {
		listInput.OrderField = *input.Query.OrderField
	}
	trainerOutputs, page, err := r.trainerService.List(&listInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	data := api_get_cms_trainers.Data{}
	if err := util.Parser(trainerOutputs, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Paging = page
	output.Data = &data
	return output
}

func (r *resolver) APIGetCMSTrainer(input *api_get_cms_trainer.Input) (output api_get_cms_trainer.Output) {
	// 查詢列表教練資訊
	findInput := model.FindInput{}
	findInput.UserID = input.Uri.UserID
	findInput.Preloads = []*preloadModel.Preload{
		{Field: "Certificates"},
		{Field: "Card"},
		{Field: "BankAccount"},
		{Field: "TrainerAlbums"},
	}
	trainerOutput, err := r.trainerService.Find(&findInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	data := api_get_cms_trainer.Data{}
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateCMSTrainer(ctx *gin.Context, tx *gorm.DB, input *api_update_cms_trainer.Input) (output api_update_cms_trainer.Output) {
	defer tx.Rollback()
	// 查詢教練資訊
	findTrainerInput := model.FindInput{}
	findTrainerInput.UserID = util.PointerInt64(input.Uri.UserID)
	findTrainerInput.Preloads = []*preload.Preload{
		{Field: "User"},
	}
	trainerOutput, err := r.trainerService.Tx(tx).Find(&findTrainerInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改教練資訊
	trainerTable := model.Table{}
	trainerTable.UserID = util.PointerInt64(input.Uri.UserID)
	trainerTable.TrainerStatus = input.Body.TrainerStatus
	trainerTable.TrainerLevel = input.Body.TrainerLevel
	if err := r.trainerService.Tx(tx).Update(&trainerTable); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// 修改教練狀態後續流程
	if input.Body.TrainerStatus != nil {
		// 添加log紀錄
		logTable := trainerStatusLogModel.Table{}
		logTable.UserID = util.PointerInt64(input.Uri.UserID)
		logTable.TrainerStatus = input.Body.TrainerStatus
		logTable.Comment = util.PointerString("")
		if _, err := r.trainerStatusLogService.Tx(tx).Create(&logTable); err != nil {
			output.Set(code.BadRequest, err.Error())
			return output
		}
		// 下架課表
		if util.OnNilJustReturnInt(input.Body.TrainerStatus, 0) == model.Suspend {
			// 查詢該用戶上架的付費課表
			courseListInput := courseModel.ListInput{}
			courseListInput.UserID = util.PointerInt64(input.Uri.UserID)
			courseListInput.SaleType = util.PointerInt(courseModel.SaleTypeCharge)
			courseListInput.CourseStatus = util.PointerInt(courseModel.Sale)
			courseOutputs, _, err := r.courseService.Tx(tx).List(&courseListInput)
			if err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
			// 將付費課表改為下架狀態
			courseTables := make([]*courseModel.Table, 0)
			for _, courseOutput := range courseOutputs {
				courseTable := courseModel.Table{}
				courseTable.ID = courseOutput.ID
				courseTable.CourseStatus = util.PointerInt(courseModel.Remove)
				courseTables = append(courseTables, &courseTable)
			}
			if err := r.courseService.Tx(tx).Updates(courseTables); err != nil {
				output.Set(code.BadRequest, err.Error())
				return output
			}
		}
	}
	tx.Commit()
	// 發送推播
	oldStatus := util.OnNilJustReturnInt(trainerOutput.TrainerStatus, 0)
	currentStatus := util.OnNilJustReturnInt(input.Body.TrainerStatus, 0)
	deviceToken := util.OnNilJustReturnString(trainerOutput.UserOnSafe().DeviceToken, "")
	if oldStatus == model.Reviewing && currentStatus == model.Activity && len(deviceToken) > 0 {
		// 準備推播訊息
		trainerName := util.OnNilJustReturnString(trainerOutput.Nickname, "")
		title := "教練審核通知"
		body := fmt.Sprintf("%v教練，你申請成為平台教練的審核已經通過囉！點此打開Fitopia.hub APP開始創建你的課表～", trainerName)
		msgOutput := fcmModel.Output{}
		msgOutput.Message.Token = deviceToken
		msgOutput.Message.Notification.Title = title
		msgOutput.Message.Notification.Body = body
		msgOutput.Message.Data.Title = title
		msgOutput.Message.Data.Body = body
		// 獲取或更新 API token
		apiToken, _ := r.redisTool.Get(r.fcmTool.Key())
		if len(apiToken) == 0 {
			//產出 auth token
			oauthToken, _ := r.fcmTool.GenerateGoogleOAuth2Token(time.Hour)
			//獲取API Token
			apiToken, _ = r.fcmTool.APIGetGooglePlayToken(oauthToken)
			//儲存API Token
			_ = r.redisTool.SetEX(r.fcmTool.Key(), apiToken, r.fcmTool.GetExpire())
		}
		message := make(map[string]interface{})
		if err = util.Parser(msgOutput, &message); err != nil {
			logger.Shared().Error(ctx, "APIUpdateCMSTrainer Parser："+err.Error())
		}
		if err = r.fcmTool.APISendMessage(apiToken, message); err != nil {
			logger.Shared().Error(ctx, "APIUpdateCMSTrainer SendMessage："+err.Error())
		}
	}
	// 查詢修改後的教練資訊
	findTrainerInput = model.FindInput{}
	findTrainerInput.UserID = util.PointerInt64(input.Uri.UserID)
	findTrainerInput.Preloads = []*preload.Preload{
		{Field: "Certificates"},
		{Field: "TrainerStatistic"},
	}
	trainerOutput, err = r.trainerService.Find(&findTrainerInput)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	// Parse Output
	data := api_update_cms_trainer.Data{}
	if err := util.Parser(trainerOutput, &data); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.Set(code.Success, "success")
	output.Data = &data
	return output
}

func (r *resolver) APIUpdateCMSTrainerAvatar(input *model.APIUpdateCMSTrainerAvatarInput) (output model.APIUpdateCMSTrainerAvatarOutput) {
	fileNamed, err := r.avatarUploadTool.Save(input.File, input.CoverNamed)
	if err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	table := model.Table{}
	table.UserID = util.PointerInt64(input.UserID)
	table.Avatar = util.PointerString(fileNamed)
	if err := r.trainerService.Update(&table); err != nil {
		output.Set(code.BadRequest, err.Error())
		return output
	}
	output.SetStatus(code.Success)
	output.Data = util.PointerString(fileNamed)
	return output
}
