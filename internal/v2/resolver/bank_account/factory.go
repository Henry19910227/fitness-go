package bank_account

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/bank_account"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	bankAccountService := bank_account.NewService(db)
	uploadTool := uploader.NewAccountImageTool()
	return New(bankAccountService, uploadTool)
}
