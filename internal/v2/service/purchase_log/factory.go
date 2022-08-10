package purchase_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/purchase_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := purchase_log.New(db)
	return New(repository)
}
