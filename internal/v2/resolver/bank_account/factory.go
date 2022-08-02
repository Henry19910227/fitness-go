package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bankAccountService := bank_account.NewService(db)
	return New(bankAccountService)
}
