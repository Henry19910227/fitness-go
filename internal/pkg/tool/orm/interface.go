package orm

import "gorm.io/gorm"

type Tool interface {
	DB() *gorm.DB
}
