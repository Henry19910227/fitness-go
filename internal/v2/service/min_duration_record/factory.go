package min_duration_record

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/min_duration_record"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := min_duration_record.New(db)
	return New(repository)
}
