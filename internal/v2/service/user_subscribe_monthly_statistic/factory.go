package user_subscribe_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_subscribe_monthly_statistic"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := user_subscribe_monthly_statistic.New(db)
	return New(repository)
}
