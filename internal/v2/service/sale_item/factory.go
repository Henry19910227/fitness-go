package sale_item

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/sale_item"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := sale_item.New(db)
	return New(repository)
}

