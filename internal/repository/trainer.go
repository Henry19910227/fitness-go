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

func (t *trainer) CreateTrainer(uid int64, param *model.CreateTrainerParam) error {
	trainer := model.Trainer{
		UserID: uid,
		Name: param.Name,
		Nickname: param.Nickname,
		Phone: param.Phone,
		Email: param.Email,
		TrainerStatus: 1,
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
		Take(entity).Error; err != nil {
		return err
	}
	return nil
}