package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type trainer struct {
	gorm tool.Gorm
}

func NewTrainer(gormTool tool.Gorm) Trainer {
	return &trainer{gorm: gormTool}
}

func (t *trainer) CreateTrainer(uid int64, param *model.CreateTrainerParam) error {
	// 創建 Trainer model
	trainer := entity.Trainer{
		UserID:        uid,
		Name:          param.Name,
		Nickname:      param.Nickname,
		Skill:         param.Skill,
		Avatar:        param.Avatar,
		Email:         param.Email,
		Phone:         param.Phone,
		Address:       param.Address,
		Intro:         param.Intro,
		Experience:    param.Experience,
		TrainerStatus: 2,
		TrainerLevel:  1,
		CreateAt:      time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:      time.Now().Format("2006-01-02 15:04:05"),
	}
	if param.Motto != nil {
		trainer.Motto = *param.Motto
	}
	if param.FacebookURL != nil {
		trainer.FacebookURL = *param.FacebookURL
	}
	if param.InstagramURL != nil {
		trainer.InstagramURL = *param.InstagramURL
	}
	if param.YoutubeURL != nil {
		trainer.YoutubeURL = *param.YoutubeURL
	}
	// 創建 card model
	card := model.Card{
		UserID:     uid,
		FrontImage: param.CardFrontImage,
		BackImage:  param.CardBackImage,
		CreateAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	// 創建相簿照片model
	var albumPhotos []*model.TrainerAlbumPhoto
	for _, photoName := range param.TrainerAlbumPhotos {
		photo := model.TrainerAlbumPhoto{
			UserID:   uid,
			Photo:    photoName,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		albumPhotos = append(albumPhotos, &photo)
	}
	// 創建證照model
	var certificates []*model.Certificate
	for i, image := range param.CertificateImages {
		certificate := model.Certificate{
			UserID:   uid,
			Name:     param.CertificateNames[i],
			Image:    image,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		certificates = append(certificates, &certificate)
	}
	// 創建銀行帳戶model
	bankAccount := model.BankAccount{
		UserID:       uid,
		AccountName:  param.AccountName,
		AccountImage: param.AccountImage,
		BankCode:     param.BankCode,
		Account:      param.Account,
		Branch:       param.Branch,
		CreateAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	// 導入db
	if err := t.gorm.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&trainer).Error; err != nil {
			return err
		}
		if err := tx.Create(&card).Error; err != nil {
			return err
		}
		if err := tx.Create(&bankAccount).Error; err != nil {
			return err
		}
		if albumPhotos != nil {
			if err := tx.Create(&albumPhotos).Error; err != nil {
				return err
			}
		}
		if certificates != nil {
			if err := tx.Create(&certificates).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (t *trainer) FindTrainer(uid int64) (*model.Trainer, error) {
	var trainer model.Trainer
	if err := t.gorm.DB().
		Table("trainers").
		Preload("TrainerStatistic").
		Joins("LEFT JOIN trainer_statistics AS ts ON ts.user_id = trainers.user_id").
		Take(&trainer, "trainers.user_id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &trainer, nil
}

func (t *trainer) FindTrainerEntity(uid int64, input interface{}) error {
	if err := t.gorm.DB().
		Model(&entity.Trainer{}).
		Where("user_id = ?", uid).
		Find(input).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainer) FindTrainerEntities(input interface{}, status *global.TrainerStatus, orderBy *model.OrderBy, paging *model.PagingParam) error {
	var db *gorm.DB
	db = t.gorm.DB().Model(&entity.Trainer{})
	if status != nil {
		db = db.Where("trainer_status = ?", *status)
	}
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	if err := db.Find(input).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainer) FindTrainersCount(status *global.TrainerStatus) (int, error) {
	var count int64
	var db *gorm.DB
	db = t.gorm.DB().Model(&model.Trainer{})
	if status != nil {
		db = db.Where("trainer_status = ?", *status)
	}
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t *trainer) UpdateTrainerByUID(uid int64, param *model.UpdateTrainerParam) error {
	if param == nil {
		return nil
	}
	var selects []interface{}
	if param.Nickname != nil {
		selects = append(selects, "nickname")
	}
	if param.Avatar != nil {
		selects = append(selects, "avatar")
	}
	if param.TrainerStatus != nil {
		selects = append(selects, "trainer_status")
	}
	if param.TrainerLevel != nil {
		selects = append(selects, "trainer_level")
	}
	if param.Intro != nil {
		selects = append(selects, "intro")
	}
	if param.Experience != nil {
		selects = append(selects, "experience")
	}
	if param.Skill != nil {
		selects = append(selects, "skill")
	}
	if param.Motto != nil {
		selects = append(selects, "motto")
	}
	if param.FacebookURL != nil {
		selects = append(selects, "facebook_url")
	}
	if param.InstagramURL != nil {
		selects = append(selects, "instagram_url")
	}
	if param.YoutubeURL != nil {
		selects = append(selects, "youtube_url")
	}
	// 插入更新時間
	selects = append(selects, "update_at")
	var updateAt = time.Now().Format("2006-01-02 15:04:05")
	param.UpdateAt = &updateAt

	// 建立待新增的相簿照片array
	var createAlbumPhotos []*model.TrainerAlbumPhoto
	for _, photoName := range param.CreateAlbumPhotos {
		photo := model.TrainerAlbumPhoto{
			UserID:   uid,
			Photo:    photoName,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		createAlbumPhotos = append(createAlbumPhotos, &photo)
	}
	// 建立待更新的證照array
	var updateCertificates []*model.Certificate
	for i, cerID := range param.UpdateCerIDs {
		certificate := model.Certificate{
			ID:       cerID,
			UserID:   uid,
			Name:     param.UpdateCerNames[i],
			Image:    param.UpdateCerImages[i],
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		updateCertificates = append(updateCertificates, &certificate)
	}
	// 建立待新增的證照array
	var createCertificates []*model.Certificate
	for i, image := range param.CreateCerImages {
		certificate := model.Certificate{
			UserID:   uid,
			Name:     param.CreateCerNames[i],
			Image:    image,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		createCertificates = append(createCertificates, &certificate)
	}
	// DB操作
	if err := t.gorm.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Table("trainers").
			Where("user_id = ?", uid).
			Select("", selects...).
			Updates(param).Error; err != nil {
			return err
		}
		//刪除指定教練相簿照片
		if len(param.DeleteAlbumPhotosIDs) > 0 {
			if err := tx.Delete(&model.TrainerAlbumPhoto{}, param.DeleteAlbumPhotosIDs).Error; err != nil {
				return err
			}
		}
		//新增指定教練相簿照片
		if createAlbumPhotos != nil {
			if err := tx.Create(&createAlbumPhotos).Error; err != nil {
				return err
			}
		}
		//刪除指定證照照片
		if len(param.DeleteCerIDs) > 0 {
			if err := tx.Delete(&model.Certificate{}, param.DeleteCerIDs).Error; err != nil {
				return err
			}
		}
		//更新指定證照照片
		for _, item := range updateCertificates {
			if err := tx.Table("certificates").
				Where("user_id = ?", uid).
				Select("name", "image", "update_at").
				Updates(item).Error; err != nil {
				return err
			}
		}
		//新增指定證照照片
		if createCertificates != nil {
			if err := tx.Create(&createCertificates).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (t *trainer) FindTrainers(result interface{}, param *model.FinsTrainersParam, orderBy *model.OrderBy, paging *model.PagingParam) (int, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 userID 篩選條件
	if param.UserID != nil {
		query += "AND user_id = ? "
		params = append(params, *param.UserID)
	}
	//加入 nickname 篩選條件
	if param.NickName != nil {
		query += "AND nickname LIKE ? "
		params = append(params, "%"+*param.NickName+"%")
	}
	//加入 email 篩選條件
	if param.Email != nil {
		query += "AND email LIKE ? "
		params = append(params, "%"+*param.Email+"%")
	}
	//加入 trainer_status 篩選條件
	if param.TrainerStatus != nil {
		query += "AND trainer_status = ? "
		params = append(params, *param.TrainerStatus)
	}
	var db *gorm.DB
	var amount int64
	db = t.gorm.DB().Model(&entity.Trainer{}).Where(query, params...).Count(&amount)

	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//分頁
	if paging != nil {
		db = db.Offset(paging.Offset).Limit(paging.Limit)
	}
	//查詢數據
	if err := db.Find(result).Error; err != nil {
		return 0, nil
	}
	return int(amount), nil
}

func (t *trainer) FindTrainerDetail(userID int64, result interface{}) error {
	if err := t.gorm.DB().
		Model(&model.TrainerDetail{}).
		Preload("BankAccount").
		Preload("TrainerAlbumPhotos").
		Preload("Certificates").
		Preload("Cards").
		Where("user_id = ?", userID).
		Take(result).Error; err != nil {
		return err
	}
	return nil
}
