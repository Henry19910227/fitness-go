package course_category_training_monthly_statistic

import "gorm.io/gorm"

func NewResolver(db *gorm.DB) Resolver {
	return New()
}
