package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type bankAccount struct {
	gorm tool.Gorm
}

func NewBankAccount(gorm tool.Gorm) BankAccount {
	return &bankAccount{gorm: gorm}
}

func (b *bankAccount) FindBankAccountEntity(userID int64, inputModel interface{}) error {
	if err := b.gorm.DB().
		Model(&entity.BankAccount{}).
		Where("user_id = ?", userID).
		Take(inputModel).Error; err != nil {
		return err
	}
	return nil
}
