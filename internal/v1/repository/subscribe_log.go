package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type subscribeLog struct {
	gorm tool.Gorm
}

func NewSubscribeLog(gorm tool.Gorm) SubscribeLog {
	return &subscribeLog{gorm: gorm}
}

func (s *subscribeLog) SaveSubscribeLog(tx *gorm.DB, param *model.CreateSubscribeLogParam) (int64, error) {
	db := s.gorm.DB()
	if tx != nil {
		db = tx
	}
	log := entity.SubscribeLog{
		OriginalTransactionID: param.OriginalTransactionID,
		TransactionID:         param.TransactionID,
		PurchaseDate:          param.PurchaseDate,
		ExpiresDate:           param.ExpiresDate,
		Type:                  param.Type,
		Msg:                   param.Msg,
		CreateAt:              time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "original_transaction_id"}, {Name: "transaction_id"}, {Name: "type"}},
		DoNothing: true,
	}).Create(&log).Error; err != nil {
		return 0, err
	}
	return log.ID, nil
}
