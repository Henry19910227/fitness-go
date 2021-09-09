package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type trainer struct {
	gorm  tool.Gorm
}

func NewTrainer(gormTool  tool.Gorm) Trainer {
	return &trainer{gorm: gormTool}
}

func (t *trainer) CreateTrainer(uid int64) error {
	trainer := model.Trainer{
		UserID: uid,
		TrainerStatus: 4,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := t.gorm.DB().Create(&trainer).Error; err != nil {
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