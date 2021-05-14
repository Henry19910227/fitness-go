package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type trainer struct {
	gorm  tool.Gorm
}

func NewTrainer(gormTool  tool.Gorm) Trainer {
	return &trainer{gorm: gormTool}
}

func (t *trainer) CreateTrainer(name string, nickname string, phone string, email string) (int64, error) {
	trainer := model.Trainer{
		Name: name,
		Nickname: nickname,
		Phone: phone,
		Email: email,
	}
	if err := t.gorm.DB().Create(&trainer).Error; err != nil {
		return 0, err
	}
	return trainer.ID, nil
}
