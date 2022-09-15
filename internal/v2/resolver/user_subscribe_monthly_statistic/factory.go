package user_subscribe_monthly_statistic

import (
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	return New()
}
