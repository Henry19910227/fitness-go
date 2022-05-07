package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type card struct {
	gorm tool.Gorm
}

func NewCard(gorm tool.Gorm) Card {
	return &card{gorm: gorm}
}

func (c *card) FindCardEntity(userID int64, inputModel interface{}) error {
	if err := c.gorm.DB().
		Model(&entity.Card{}).
		Where("user_id = ?", userID).
		Take(inputModel).Error; err != nil {
		return err
	}
	return nil
}
