package course_usage_statistic

import "gorm.io/gorm"

type Service interface {
	Tx(tx *gorm.DB) Service
	Statistic() (err error)
}
