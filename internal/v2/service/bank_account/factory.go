package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/bank_account"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := bank_account.New(db)
	return New(repository)
}
