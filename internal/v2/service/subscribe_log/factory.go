package subscribe_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/subscribe_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := subscribe_log.New(db)
	return New(repository)
}
