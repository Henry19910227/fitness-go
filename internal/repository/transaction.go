package repository

import (
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
)

type transaction struct {
	gorm tool.Gorm
}

func NewTransaction(gorm tool.Gorm) Transaction {
	return &transaction{gorm: gorm}
}

func (t *transaction) CreateTransaction() *gorm.DB {
	return t.gorm.DB().Begin()
}

func (t *transaction) FinishTransaction(tx *gorm.DB) {
	tx.Commit()
	tx = nil
}

