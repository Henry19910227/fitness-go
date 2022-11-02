package trainer_album

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer_album"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := trainer_album.New(db)
	return New(repository)
}

