package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/bank_account"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := bank_account.NewResolver(db)
	return New(resolver)
}
