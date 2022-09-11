package max_distance_record

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/max_distance_record"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := max_distance_record.New(db)
	return New(repository)
}