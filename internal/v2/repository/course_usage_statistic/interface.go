package course_usage_statistic

import (
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	Statistic() (err error)
}
