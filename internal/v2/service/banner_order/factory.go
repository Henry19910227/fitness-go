package banner_order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/banner_order"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := banner_order.New(db)
	return New(repository)
}
