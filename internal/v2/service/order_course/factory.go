package order_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/order_course"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := order_course.New(db)
	return New(repository)
}