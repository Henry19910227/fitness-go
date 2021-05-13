package repository

import "github.com/Henry19910227/fitness-go/internal/tool"

type trainer struct {
	gorm  tool.Gorm
}

func NewTrainer(gormTool  tool.Gorm) Trainer {
	return &trainer{gorm: gormTool}
}

func (t *trainer) CreateTrainer(name string, nickname string, phone string, email string) (int64, error) {
	panic("implement me")
}
