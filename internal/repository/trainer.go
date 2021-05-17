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

func (t *trainer) CreateTrainer(uid int64, param *model.CreateTrainerParam) (int64, error) {
	trainer := model.Trainer{
		Name: param.Name,
		Nickname: param.Nickname,
		Phone: param.Phone,
		Email: param.Email,
		UserID: uid,
		TrainerStatus: 1,
		Birthday: "0000-01-01 00:00:00",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := t.gorm.DB().Create(&trainer).Error; err != nil {
		return 0, err
	}
	return trainer.ID, nil
}
