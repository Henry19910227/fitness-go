package trainer

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/code"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	accountModel "github.com/Henry19910227/fitness-go/internal/v2/model/bank_account"
	cardModel "github.com/Henry19910227/fitness-go/internal/v2/model/card"
	certModel "github.com/Henry19910227/fitness-go/internal/v2/model/certificate"
	joinModel "github.com/Henry19910227/fitness-go/internal/v2/model/join"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	orderByModel "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	albumModel "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_album"
	whereModel "github.com/Henry19910227/fitness-go/internal/v2/model/where"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
	"github.com/Henry19910227/fitness-go/internal/v2/service/card"
	"github.com/Henry19910227/fitness-go/internal/v2/service/certificate"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer"
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_album"
	"gorm.io/gorm"
)

type resolver struct {
	trainerService trainer.Service
	trainerAlbumService trainer_album.Service
	cardService card.Service
	certService certificate.Service
	bankAccountService bank_account.Service
	avatarUploadTool uploader.Tool
	albumUploadTool uploader.Tool
	cardFrontUploadTool uploader.Tool
	cardBackUploadTool uploader.Tool
	certUploadTool    uploader.Tool
	accountUploadTool uploader.Tool
}

func New(trainerService trainer.Service, trainerAlbumService trainer_album.Service,
	cardService card.Service, certService certificate.Service,
	bankAccountService bank_account.Service,
	avatarUploadTool uploader.Tool, albumUploadTool uploader.Tool,
	cardFrontUploadTool uploader.Tool, cardBackUploadTool uploader.Tool,
	certUploadTool uploader.Tool, accountUploadTool uploader.Tool) Resolver {
	return &resolver{trainerService: trainerService, trainerAlbumService: trainerAlbumService,
		cardService: cardService, certService: certService,
		bankAccountService: bankAccountService,
		avatarUploadTool: avatarUploadTool, albumUploadTool: albumUploadTool,
		cardFrontUploadTool: cardFrontUploadTool, cardBackUploadTool: cardBackUploadTool,
		certUploadTool: certUploadTool, accountUploadTool: accountUploadTool}
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
	trainerTable.TrainerLevel = util.PointerInt(0)
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

func (r *resolver) APIGetTrainerProfile(input *model.APIGetTrainerProfileInput) (output model.APIGetTrainerProfileOutput) {
	listInput := model.ListInput{}
	listInput.UserID = util.PointerInt64(input.UserID)
	listInput.Preloads = []*preload.Preload{{Field: "TrainerStatistic"}, {Field: "Certificates"}, {Field: "TrainerAlbums"}}
	listInput.Size = 1
	listInput.Page = 1
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
