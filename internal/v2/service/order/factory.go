package order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/order"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := order.New(db)
	return New(repository)
}
