package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type trainer struct {
	gorm  tool.Gorm
}

func NewTrainer(gormTool  tool.Gorm) Trainer {
	return &trainer{gorm: gormTool}
}

func (t *trainer) CreateTrainer(uid int64, param *model.CreateTrainerParam) error {
	// 創建 Trainer model
	trainer := model.Trainer{
		UserID: uid,
		Name: param.Name,
		Nickname: param.Nickname,
		Skill: param.Skill,
		Avatar: param.Avatar,
		Email: param.Email,
		Phone: param.Phone,
		Address: param.Address,
		Intro: param.Intro,
		Experience: param.Experience,
		TrainerStatus: 2,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
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
		UserID: uid,
		FrontImage: param.CardFrontImage,
		BackImage: param.CardBackImage,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 創建相簿照片model
	var albumPhotos []*model.TrainerAlbumPhoto
	for _, photoName := range param.TrainerAlbumPhotos {
		 photo := model.TrainerAlbumPhoto{
		 	UserID: uid,
			Photo: photoName,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		 }
		albumPhotos = append(albumPhotos, &photo)
	}
	// 創建證照model
	var certificates []*model.Certificate
	for i, image := range param.CertificateImages {
		certificate := model.Certificate{
			UserID: uid,
			Name: param.CertificateNames[i],
			Image: image,
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
		BackCode:     param.BankCode,
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

func (t *trainer) FindTrainerByUID(uid int64, entity interface{}) error {
	if err := t.gorm.DB().
		Model(&model.Trainer{}).
		Where("user_id = ?", uid).
		Find(entity).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainer) UpdateTrainerByUID(uid int64, param *model.UpdateTrainerParam) error {
	var selects []interface{}
	if param.Name != nil { selects = append(selects, "name") }
	if param.Nickname != nil { selects = append(selects, "nickname") }
	if param.Avatar != nil { selects = append(selects, "avatar") }
	if param.TrainerStatus != nil { selects = append(selects, "trainer_status") }
	if param.Email != nil { selects = append(selects, "email") }
	if param.Phone != nil { selects = append(selects, "phone") }
	if param.Address != nil { selects = append(selects, "address") }
	if param.Intro != nil { selects = append(selects, "intro") }
	if param.Experience != nil { selects = append(selects, "experience") }
	if param.Motto != nil { selects = append(selects, "motto") }
	if param.CardID != nil { selects = append(selects, "card_id") }
	if param.CardFrontImage != nil { selects = append(selects, "card_front_image") }
	if param.CardBackImage != nil { selects = append(selects, "card_back_image") }
	if param.FacebookURL != nil { selects = append(selects, "facebook_url") }
	if param.InstagramURL != nil { selects = append(selects, "instagram_url") }
	if param.YoutubeURL != nil { selects = append(selects, "youtube_url") }
	//插入更新時間
	if param != nil {
		selects = append(selects, "update_at")
		var updateAt = time.Now().Format("2006-01-02 15:04:05")
		param.UpdateAt = &updateAt
	}
	if err := t.gorm.DB().
		Table("trainers").
		Where("user_id = ?", uid).
		Select("", selects...).
		Updates(param).Error; err != nil {
		return err
	}
	return nil
}