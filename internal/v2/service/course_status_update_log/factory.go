package course_status_update_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course_status_update_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := course_status_update_log.New(db)
	return New(repository)
}
