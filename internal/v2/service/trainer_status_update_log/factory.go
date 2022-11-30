package trainer_status_update_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer_status_update_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := trainer_status_update_log.New(db)
	return New(repository)
}
